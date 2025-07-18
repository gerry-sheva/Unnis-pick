package service

import (
	"context"
	"unnis_pick/internal/database/mapper"
	"unnis_pick/internal/database/query"
	"unnis_pick/internal/domain"
)

type ProductService struct {
	queries *query.Queries
}

func NewProductService(queries *query.Queries) *ProductService {
	return &ProductService{
		queries: queries,
	}
}

func (s *ProductService) CreateProduct(ctx context.Context, product *domain.Product) (*domain.Product, error) {
	uuid, err := mapper.StringToUUID(product.BrandID)
	if err != nil {
		return nil, err
	}

	price, err := mapper.Float64ToNumeric(product.Price)
	if err != nil {
		return nil, err
	}

	params := query.CreateProductParams{
		Name:    product.Name,
		Price:   price,
		Stock:   product.Stock,
		BrandID: uuid,
	}
	dbModel, err := s.queries.CreateProduct(ctx, params)
	if err != nil {
		return nil, err
	}
	product, err = mapper.DBModelToProduct(&dbModel)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (s *ProductService) GetProduct(ctx context.Context, productID string) (*domain.Product, error) {
	uuid, err := mapper.StringToUUID(productID)
	if err != nil {
		return nil, err
	}
	dbModel, err := s.queries.GetProduct(ctx, uuid)
	if err != nil {
		return nil, err
	}
	product, err := mapper.DBModelToProduct(&dbModel)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (s *ProductService) UpdateProduct(ctx context.Context, productId string, product *domain.Product) (*domain.Product, error) {
	uuid, err := mapper.StringToUUID(productId)
	if err != nil {
		return nil, err
	}

	brandId, err := mapper.StringToUUID(product.BrandID)
	if err != nil {
		return nil, err
	}

	price, err := mapper.Float64ToNumeric(product.Price)
	if err != nil {
		return nil, err
	}

	params := query.UpdateProductParams{
		ProductID: uuid,
		Name:      product.Name,
		Price:     price,
		Stock:     product.Stock,
		BrandID:   brandId,
	}
	dbModel, err := s.queries.UpdateProduct(ctx, params)
	if err != nil {
		return nil, err
	}
	product, err = mapper.DBModelToProduct(&dbModel)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (s *ProductService) DeleteProduct(ctx context.Context, productID string) error {
	uuid, err := mapper.StringToUUID(productID)
	if err != nil {
		return err
	}
	err = s.queries.DeleteProduct(ctx, uuid)
	if err != nil {
		return err
	}
	return nil
}
