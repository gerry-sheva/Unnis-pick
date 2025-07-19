package service

import (
	"context"
	"log"
	"unnis_pick/internal/database/mapper"
	"unnis_pick/internal/database/query"
	"unnis_pick/internal/domain"
	"unnis_pick/internal/searchengine"

	"github.com/elastic/go-elasticsearch/v9/typedapi/core/search"
)

type ProductService struct {
	queries       *query.Queries
	searchService searchengine.Service
}

func NewProductService(queries *query.Queries, searchService searchengine.Service) *ProductService {
	return &ProductService{
		queries:       queries,
		searchService: searchService,
	}
}

func (s *ProductService) CreateProduct(ctx context.Context, product *domain.Product) (*domain.Product, error) {
	brandId, err := mapper.StringToUUID(product.BrandID)
	if err != nil {
		return nil, err
	}

	price, err := mapper.Float64ToNumeric(product.Price)
	if err != nil {
		log.Println("this here")
		return nil, err
	}

	params := query.CreateProductParams{
		Name:        product.Name,
		Description: product.Description,
		Price:       price,
		Stock:       product.Stock,
		BrandID:     brandId,
	}
	dbModel, err := s.queries.CreateProduct(ctx, params)
	if err != nil {
		return nil, err
	}

	product, err = mapper.DBModelToProduct(&dbModel)
	if err != nil {
		return nil, err
	}

	brand, err := s.queries.GetBrand(ctx, brandId)
	if err != nil {
		return nil, err
	}

	err = s.searchService.IndexProduct(product, brand.Name)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (s *ProductService) GetProduct(ctx context.Context, id string) (*domain.Product, error) {
	productId, err := mapper.StringToUUID(id)
	if err != nil {
		return nil, err
	}
	dbModel, err := s.queries.GetProduct(ctx, productId)
	if err != nil {
		return nil, err
	}
	product, err := mapper.DBModelToProduct(&dbModel)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (s *ProductService) UpdateProduct(ctx context.Context, id string, product *domain.Product) (*domain.Product, error) {
	productId, err := mapper.StringToUUID(id)
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
		ProductID:   productId,
		Name:        product.Name,
		Description: product.Description,
		Price:       price,
		Stock:       product.Stock,
		BrandID:     brandId,
	}
	dbModel, err := s.queries.UpdateProduct(ctx, params)
	if err != nil {
		return nil, err
	}
	product, err = mapper.DBModelToProduct(&dbModel)
	if err != nil {
		return nil, err
	}

	brand, err := s.queries.GetBrand(ctx, brandId)
	if err != nil {
		return nil, err
	}

	err = s.searchService.IndexProduct(product, brand.Name)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (s *ProductService) DeleteProduct(ctx context.Context, id string) error {
	productId, err := mapper.StringToUUID(id)
	if err != nil {
		return err
	}
	err = s.queries.DeleteProduct(ctx, productId)
	if err != nil {
		return err
	}

	s.searchService.DeleteProduct(id)

	return nil
}

func (s *ProductService) QueryProducts(ctx context.Context, filter *domain.ProductFilter) (*search.Response, error) {
	res := s.searchService.QueryProducts(filter)

	return res, nil
}
