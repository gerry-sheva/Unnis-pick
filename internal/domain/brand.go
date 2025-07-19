package domain

import (
	"context"
	"time"
)

type Brand struct {
	BrandID   string     `json:"id"`
	Name      string     `json:"name"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

type BrandService interface {
	CreateBrand(ctx context.Context, brand *Brand) (*Brand, error)
	GetBrand(ctx context.Context, id string) (*Brand, error)
	UpdateBrand(ctx context.Context, id string, brand *Brand) (*Brand, error)
	DeleteBrand(ctx context.Context, id string) error
}
