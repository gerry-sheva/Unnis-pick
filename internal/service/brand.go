package service

import (
	"context"
	"unnis_pick/internal/database/mapper"
	"unnis_pick/internal/database/query"
	"unnis_pick/internal/domain"
)

type BrandService struct {
	queries *query.Queries
}

func NewBrandService(queries *query.Queries) *BrandService {
	return &BrandService{
		queries: queries,
	}
}

func (s *BrandService) CreateBrand(ctx context.Context, brand *domain.Brand) (*domain.Brand, error) {
	dbBrand, err := s.queries.CreateBrand(ctx, brand.Name)
	if err != nil {
		return nil, err
	}
	brand, err = mapper.DBModelToBrand(&dbBrand)
	if err != nil {
		return nil, err
	}
	return brand, nil
}

func (s *BrandService) GetBrand(ctx context.Context, brandID string) (*domain.Brand, error) {
	uuid, err := mapper.StringToUUID(brandID)
	if err != nil {
		return nil, err
	}

	dbBrand, err := s.queries.GetBrand(ctx, uuid)
	if err != nil {
		return nil, err
	}
	brand, err := mapper.DBModelToBrand(&dbBrand)
	if err != nil {
		return nil, err
	}
	return brand, nil
}

func (s *BrandService) UpdateBrand(ctx context.Context, brandId string, brand *domain.Brand) (*domain.Brand, error) {
	uuid, err := mapper.StringToUUID(brandId)
	if err != nil {
		return nil, err
	}

	params := query.UpdateBrandParams{
		BrandID: uuid,
		Name:    brand.Name,
	}
	dbBrand, err := s.queries.UpdateBrand(ctx, params)
	if err != nil {
		return nil, err
	}

	brand, err = mapper.DBModelToBrand(&dbBrand)
	if err != nil {
		return nil, err
	}

	return brand, nil
}

func (s *BrandService) DeleteBrand(ctx context.Context, brandID string) error {
	uuid, err := mapper.StringToUUID(brandID)
	if err != nil {
		return err
	}

	err = s.queries.DeleteBrand(ctx, uuid)
	if err != nil {
		return err
	}

	return nil
}
