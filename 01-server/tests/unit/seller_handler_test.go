package unit

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"01-server/internal/handlers"
	"01-server/internal/models"
	"01-server/tests/mocks"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetAllSellers_Success(t *testing.T) {
	mockRepo := new(mocks.SellerRepositoryMock)
	handler := handlers.NewSellerHandler(mockRepo)

	expectedSellers := []models.Seller{
		{ID: 1, Name: "Seller One", Phone: "123456789"},
		{ID: 2, Name: "Seller Two", Phone: "987654321"},
	}

	mockRepo.On("GetAll").Return(expectedSellers, nil)

	req, _ := http.NewRequest("GET", "/sellers", nil)
	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	var sellers []models.Seller
	json.NewDecoder(rr.Body).Decode(&sellers)
	assert.Equal(t, expectedSellers, sellers)

	mockRepo.AssertExpectations(t)
}

func TestGetSellerByID_Success(t *testing.T) {
	mockRepo := new(mocks.SellerRepositoryMock)
	handler := handlers.NewSellerHandler(mockRepo)

	expectedSeller := models.Seller{ID: 1, Name: "Seller One", Phone: "123456789"}
	mockRepo.On("GetByID", 1).Return(expectedSeller, nil)

	req, _ := http.NewRequest("GET", "/sellers/1", nil)
	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	var seller models.Seller
	json.NewDecoder(rr.Body).Decode(&seller)
	assert.Equal(t, expectedSeller, seller)

	mockRepo.AssertExpectations(t)
}

func TestCreateSeller_Success(t *testing.T) {
	mockRepo := new(mocks.SellerRepositoryMock)
	handler := handlers.NewSellerHandler(mockRepo)

	newSeller := models.Seller{Name: "New Seller", Phone: "111222333"}
	mockRepo.On("Create", mock.Anything).Return(1, nil)

	body, _ := json.Marshal(newSeller)
	req, _ := http.NewRequest("POST", "/sellers", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code)

	mockRepo.AssertExpectations(t)
}

func TestUpdateSeller_Success(t *testing.T) {
	mockRepo := new(mocks.SellerRepositoryMock)
	handler := handlers.NewSellerHandler(mockRepo)

	updatedSeller := models.Seller{ID: 1, Name: "Updated Seller", Phone: "555666777"}
	mockRepo.On("Update", updatedSeller).Return(nil)

	body, _ := json.Marshal(updatedSeller)
	req, _ := http.NewRequest("PUT", "/sellers/1", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	mockRepo.AssertExpectations(t)
}

func TestDeleteSeller_Success(t *testing.T) {
	mockRepo := new(mocks.SellerRepositoryMock)
	handler := handlers.NewSellerHandler(mockRepo)

	mockRepo.On("Delete", 1).Return(nil)

	req, _ := http.NewRequest("DELETE", "/sellers/1", nil)
	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusNoContent, rr.Code)

	mockRepo.AssertExpectations(t)
}
