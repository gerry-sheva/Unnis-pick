package mapper

import (
	"unnis_pick/internal/database/query"
	"unnis_pick/internal/domain"
)

func DBModelToProduct(dbModel *query.Product) (*domain.Product, error) {
	price, err := dbModel.Price.Float64Value()
	if err != nil {
		return nil, err
	}
	return &domain.Product{
		ProductID:   dbModel.ProductID.String(),
		Name:        dbModel.Name,
		Description: dbModel.Description,
		Price:       price.Float64,
		Stock:       int32(dbModel.Stock),
		CreatedAt:   dbModel.CreatedAt.Time,
		UpdatedAt:   &dbModel.UpdatedAt.Time,
		DeletedAt:   &dbModel.DeletedAt.Time,
		BrandID:     dbModel.BrandID.String(),
	}, nil
}
