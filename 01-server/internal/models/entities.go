package models

// Seller
type Seller struct {
	ID    int    `json:"id,omitempty"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

type SellerRequest struct {
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

// Product
type Product struct {
	ID          int     `json:"id,omitempty"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	SellerID    int     `json:"seller_id"`
}

// Customer
type Customer struct {
	ID    int    `json:"id,omitempty"`
	Name  string `json:"name"`
	Phone string `json:"phone"`
}

// Order
type Order struct {
	ID         int   `json:"id,omitempty"`
	CustomerID int   `json:"customer_id"`
	Products   []int `json:"products"` // list of product IDs
}
