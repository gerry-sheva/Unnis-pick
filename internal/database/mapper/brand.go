package mapper

import (
	"unnis_pick/internal/database/query"
	"unnis_pick/internal/domain"
)

func DBModelToBrand(dbBrand *query.Brand) (*domain.Brand, error) {
	return &domain.Brand{
		BrandID:   dbBrand.BrandID.String(),
		Name:      dbBrand.Name,
		CreatedAt: dbBrand.CreatedAt.Time,
		UpdatedAt: &dbBrand.UpdatedAt.Time,
		DeletedAt: &dbBrand.DeletedAt.Time,
	}, nil
}
