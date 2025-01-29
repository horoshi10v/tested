package mocks

import (
	"01-server/internal/models"
	"01-server/internal/repository"

	"github.com/stretchr/testify/mock"
)

var _ repository.UserRepository = (*MockUserRepo)(nil)

type MockUserRepo struct {
	mock.Mock
}

func (m *MockUserRepo) CreateUser(user models.User) (int, error) {
	args := m.Called(user)
	return args.Int(0), args.Error(1)
}

func (m *MockUserRepo) GetByUsername(username string) (models.User, error) {
	args := m.Called(username)
	return args.Get(0).(models.User), args.Error(1)
}
