package domain

import (
	"context"
	"encoding/json"
	"errors"
	"time"
)

var maxPrice = 100_000_000.0

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
	QueryProducts(ctx context.Context, filter *ProductFilter) (*ProductQueryResponse, error)
}

type ProductFilter struct {
	Name       string  `query:"name"`
	Brand      string  `query:"brand"`
	PriceMin   float64 `query:"price_min"`
	PriceMax   float64 `query:"price_max"`
	PageSize   int     `query:"size"`
	PageNumber int     `query:"number"`
}

type ProductQuery struct {
	ProductID   string  `json:"product_id"`
	Brand       string  `json:"brand"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Stock       int32   `json:"stock"`
}

type ProductQueryResponse struct {
	Total int64          `json:"total"`
	Items []ProductQuery `json:"items"`
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
	if f.PriceMin != 0 && f.PriceMax == 0 {
		f.PriceMax = maxPrice
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

func (r *ProductQueryResponse) UnmarshalJSON(data []byte) error {
	var raw struct {
		Hits struct {
			Total struct {
				Value int64 `json:"value"`
			} `json:"total"`
			Hits []struct {
				Source ProductQuery `json:"_source"`
			} `json:"hits"`
		} `json:"hits"`
	}

	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}

	r.Total = raw.Hits.Total.Value
	r.Items = make([]ProductQuery, len(raw.Hits.Hits))
	for i, hit := range raw.Hits.Hits {
		r.Items[i] = hit.Source
	}

	return nil
}
