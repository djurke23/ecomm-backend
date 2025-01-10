package models

import "gorm.io/gorm"

// Product - model proizvoda (tabela products)
type Product struct {
	gorm.Model
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	// Ako želiš, dodaj npr. Category, ImageURL, itd.
}
