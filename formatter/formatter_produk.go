package formatter

import (
	"camera-rent/entity"
	"time"
)

type Product struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	RentCost    int       `json:"rent_cost"`
	Stock       int       `json:"stock"`
	Description string    `json:"description"`
	CategoryID  int       `json:"category_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func FormatterProduct(produk *entity.Products) Product {

	return Product{
		ID:          produk.ID,
		Name:        produk.Name,
		RentCost:    produk.RentCost,
		Stock:       produk.Stock,
		Description: produk.Description,
		CategoryID:  produk.CategoryID,
		CreatedAt:   produk.CreatedAt,
		UpdatedAt:   produk.UpdatedAt,
	}
}

func FormatterGetProducts(produk []*entity.Products) []Product {
	produkGetFormatter := []Product{}

	for _, product := range produk {
		produkFormatter := FormatterProduct(product)
		produkGetFormatter = append(produkGetFormatter, produkFormatter)
	}

	return produkGetFormatter
}
