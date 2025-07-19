package domain

import (
	"context"
	"errors"
	"time"

	"github.com/elastic/go-elasticsearch/v9/typedapi/core/search"
)

type Product struct {
	ProductID   string     `json:"id"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	Price       float64    `json:"price"`
	Stock       int32      `json:"stock"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at"`
	BrandID     string     `json:"brand_id"`
}

type ProductService interface {
	CreateProduct(ctx context.Context, product *Product) (*Product, error)
	GetProduct(ctx context.Context, id string) (*Product, error)
	UpdateProduct(ctx context.Context, id string, product *Product) (*Product, error)
	DeleteProduct(ctx context.Context, id string) error
	QueryProducts(ctx context.Context, filter *ProductFilter) (*search.Response, error)
}

type ProductFilter struct {
	Name       string  `query:"name"`
	Brand      string  `query:"brand"`
	PriceMin   float64 `query:"price_min"`
	PriceMax   float64 `query:"price_max"`
	PageSize   int     `query:"size"`
	PageNumber int     `query:"number"`
}

// SetDefault sets default values for PageSize and PageNumber fields
// of the ProductFilter struct. If PageSize is 0, it defaults to 10.
// If PageNumber is 0, it remains at 0 (first page).
func (f *ProductFilter) SetDefault() {
	if f.PageSize == 0 {
		f.PageSize = 10
	}
	if f.PageNumber == 0 {
		f.PageNumber = 0
	}
}

// Validate checks the logical validity of the ProductFilter fields.
// It ensures that PriceMin is not greater than PriceMax.
// Returns an error if the price range is invalid, otherwise returns nil.
func (f *ProductFilter) Validate() error {
	if f.PriceMin > f.PriceMax {
		return errors.New("invalid price range")
	}
	return nil
}
