package models

// Product represents a product
// @Description product model
// @receiver p
// @return Product
type Product struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price float64 `json:"price"`
}