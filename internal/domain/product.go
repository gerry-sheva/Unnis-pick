package domain

import (
	"context"
	"time"
)

type Product struct {
	ProductID string     `json:"product_id"`
	Name      string     `json:"name"`
	Price     float64    `json:"price"`
	Stock     int32      `json:"stock"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
	BrandID   string     `json:"brand_id"`
}

type ProductService interface {
	CreateProduct(ctx context.Context, product *Product) (*Product, error)
	GetProduct(ctx context.Context, productID string) (*Product, error)
	UpdateProduct(ctx context.Context, productID string, product *Product) (*Product, error)
	DeleteProduct(ctx context.Context, productID string) error
}
