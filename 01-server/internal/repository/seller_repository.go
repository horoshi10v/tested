package repository

import (
	"database/sql"

	"01-server/internal/models"
)

type SellerRepository interface {
	GetAll() ([]models.Seller, error)
	GetByID(id int) (models.Seller, error)
	Create(s models.Seller) (int, error)
	Update(s models.Seller) error
	Delete(id int) error
}

type sellerRepo struct {
	db *sql.DB
}

func NewSellerRepo(db *sql.DB) SellerRepository {
	return &sellerRepo{db: db}
}

func (r *sellerRepo) GetAll() ([]models.Seller, error) {
	rows, err := r.db.Query("SELECT id, name, phone FROM sellers")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var sellers []models.Seller
	for rows.Next() {
		var s models.Seller
		if err := rows.Scan(&s.ID, &s.Name, &s.Phone); err != nil {
			return nil, err
		}
		sellers = append(sellers, s)
	}
	return sellers, nil
}

func (r *sellerRepo) GetByID(id int) (models.Seller, error) {
	var s models.Seller
	err := r.db.QueryRow("SELECT id, name, phone FROM sellers WHERE id=$1", id).
		Scan(&s.ID, &s.Name, &s.Phone)
	return s, err
}

func (r *sellerRepo) Create(s models.Seller) (int, error) {
	var newID int
	err := r.db.QueryRow(
		"INSERT INTO sellers (name, phone) VALUES ($1, $2) RETURNING id",
		s.Name, s.Phone,
	).Scan(&newID)
	return newID, err
}

func (r *sellerRepo) Update(s models.Seller) error {
	_, err := r.db.Exec(
		"UPDATE sellers SET name=$1, phone=$2 WHERE id=$3",
		s.Name, s.Phone, s.ID,
	)
	return err
}

func (r *sellerRepo) Delete(id int) error {
	_, err := r.db.Exec("DELETE FROM sellers WHERE id=$1", id)
	return err
}
