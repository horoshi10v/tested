package repository

import (
	"database/sql"

	"01-server/internal/models"
)

type UserRepository interface {
	CreateUser(user models.User) (int, error)
	GetByUsername(username string) (models.User, error)
}

type userRepo struct {
	db *sql.DB
}

func NewUserRepo(db *sql.DB) UserRepository {
	return &userRepo{db: db}
}

func (r *userRepo) CreateUser(u models.User) (int, error) {
	var newID int
	err := r.db.QueryRow(
		"INSERT INTO users (username, password) VALUES ($1, $2) RETURNING id",
		u.Username, u.Password,
	).Scan(&newID)
	return newID, err
}

func (r *userRepo) GetByUsername(username string) (models.User, error) {
	var user models.User
	err := r.db.QueryRow(
		"SELECT id, username, password FROM users WHERE username=$1",
		username,
	).Scan(&user.ID, &user.Username, &user.Password)

	return user, err
}
