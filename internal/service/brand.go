package service

import (
	"context"
	"errors"
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

func (s *BrandService) GetBrand(ctx context.Context, id string) (*domain.Brand, error) {
	brandId, err := mapper.StringToUUID(id)
	if err != nil {
		return nil, err
	}

	dbBrand, err := s.queries.GetBrand(ctx, brandId)
	if err != nil {
		return nil, err
	}
	brand, err := mapper.DBModelToBrand(&dbBrand)
	if err != nil {
		return nil, err
	}
	return brand, nil
}

func (s *BrandService) UpdateBrand(ctx context.Context, id string, brand *domain.Brand) (*domain.Brand, error) {
	brandId, err := mapper.StringToUUID(id)
	if err != nil {
		return nil, err
	}

	params := query.UpdateBrandParams{
		BrandID: brandId,
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

func (s *BrandService) DeleteBrand(ctx context.Context, id string) error {
	brandId, err := mapper.StringToUUID(id)
	if err != nil {
		return err
	}

	exists, err := s.queries.IsBrandUsed(ctx, brandId)
	if err != nil {
		return err
	}

	if exists {
		return errors.New("Brand is currently being used!")
	}

	err = s.queries.DeleteBrand(ctx, brandId)
	if err != nil {
		return err
	}

	return nil
}
