package unit

import (
	"01-server/internal/handlers"
	"01-server/internal/models"
	"01-server/tests/mocks"
	"encoding/json"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestRegister_Success(t *testing.T) {
	mockRepo := new(mocks.MockUserRepo)
	mockRepo.On("CreateUser", mock.Anything).Return(1, nil)

	handler := handlers.NewAuthHandler(mockRepo, "test_secret")

	reqBody := `{"username": "testuser", "password": "securepass"}`
	req := httptest.NewRequest(http.MethodPost, "/register", strings.NewReader(reqBody))
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	handler.Register(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code)

	var resp map[string]string
	err := json.Unmarshal(rr.Body.Bytes(), &resp)
	assert.NoError(t, err)
	assert.Equal(t, "User created and logged in", resp["message"])

	mockRepo.AssertExpectations(t)
}

func TestLogin_Success(t *testing.T) {
	mockRepo := new(mocks.MockUserRepo)
	hashedPass, _ := bcrypt.GenerateFromPassword([]byte("securepass"), bcrypt.DefaultCost)
	mockRepo.On("GetByUsername", "testuser").Return(models.User{ID: 1, Username: "testuser", Password: string(hashedPass)}, nil)

	handler := handlers.NewAuthHandler(mockRepo, "test_secret")

	reqBody := `{"username": "testuser", "password": "securepass"}`
	req := httptest.NewRequest(http.MethodPost, "/login", strings.NewReader(reqBody))
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	handler.Login(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var resp map[string]string
	err := json.Unmarshal(rr.Body.Bytes(), &resp)
	assert.NoError(t, err)
	assert.Equal(t, "Login successful", resp["message"])

	mockRepo.AssertExpectations(t)
}
