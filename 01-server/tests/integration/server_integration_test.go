package integration

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestServerIntegration(t *testing.T) {
	serverURL := "http://localhost:8081"
	client := &http.Client{}
	var authToken string

	time.Sleep(2 * time.Second)

	log.Println("Registering test user...")
	registerBody := `{"username": "testuser", "password": "securepass"}`
	req, _ := http.NewRequest("POST", serverURL+"/register", bytes.NewBufferString(registerBody))
	req.Header.Set("Content-Type", "application/json")
	resp, err := client.Do(req)
	assert.NoError(t, err)

	body, _ := io.ReadAll(resp.Body)
	log.Printf("Response: %s\n", string(body))

	if resp.StatusCode == http.StatusConflict {
		log.Println("User already exists, logging in...")

		loginBody := `{"username": "testuser", "password": "securepass"}`
		req, _ := http.NewRequest("POST", serverURL+"/login", bytes.NewBufferString(loginBody))
		req.Header.Set("Content-Type", "application/json")
		resp, err := client.Do(req)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, resp.StatusCode)

		for _, cookie := range resp.Cookies() {
			if cookie.Name == "access_token" {
				authToken = cookie.Value
				break
			}
		}
		assert.NotEmpty(t, authToken, "JWT token should not be empty after login")

	} else {
		assert.Equal(t, http.StatusCreated, resp.StatusCode)

		for _, cookie := range resp.Cookies() {
			if cookie.Name == "access_token" {
				authToken = cookie.Value
				break
			}
		}
		assert.NotEmpty(t, authToken, "JWT token should not be empty after registration")
	}

	log.Println("Fetching sellers...")
	req, _ = http.NewRequest("GET", serverURL+"/sellers", nil)
	resp, err = client.Do(req)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)

	log.Println("Creating a new seller...")
	sellerBody := `{"name": "Test Seller", "phone": "123456789"}`
	req, _ = http.NewRequest("POST", serverURL+"/sellers", bytes.NewBufferString(sellerBody))
	req.Header.Set("Content-Type", "application/json")

	req.Header.Set("Authorization", "Bearer "+authToken)

	resp, err = client.Do(req)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusCreated, resp.StatusCode)
}
