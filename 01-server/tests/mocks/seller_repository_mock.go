package mocks

import (
	"01-server/internal/models"

	"github.com/stretchr/testify/mock"
)

type SellerRepositoryMock struct {
	mock.Mock
}

func (m *SellerRepositoryMock) GetAll() ([]models.Seller, error) {
	args := m.Called()
	return args.Get(0).([]models.Seller), args.Error(1)
}

func (m *SellerRepositoryMock) GetByID(id int) (models.Seller, error) {
	args := m.Called(id)
	return args.Get(0).(models.Seller), args.Error(1)
}

func (m *SellerRepositoryMock) Create(s models.Seller) (int, error) {
	args := m.Called(s)
	return args.Int(0), args.Error(1)
}

func (m *SellerRepositoryMock) Update(s models.Seller) error {
	args := m.Called(s)
	return args.Error(0)
}

func (m *SellerRepositoryMock) Delete(id int) error {
	args := m.Called(id)
	return args.Error(0)
}
