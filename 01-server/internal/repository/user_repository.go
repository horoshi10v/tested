package repository

import (
	"database/sql"
	"errors"
	"github.com/lib/pq"

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

	if err != nil {
		var pgErr *pq.Error
		if errors.As(err, &pgErr) && pgErr.Code == "23505" { // 23505 = duplicate key
			return 0, errors.New("user already exists")
		}
		return 0, err
	}
	return newID, nil
}

func (r *userRepo) GetByUsername(username string) (models.User, error) {
	var user models.User
	err := r.db.QueryRow(
		"SELECT id, username, password FROM users WHERE username=$1",
		username,
	).Scan(&user.ID, &user.Username, &user.Password)

	return user, err
}
