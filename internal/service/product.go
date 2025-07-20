package service

import (
	"context"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"log"
	"time"
	"unnis_pick/internal/cache"
	"unnis_pick/internal/database/mapper"
	"unnis_pick/internal/database/query"
	"unnis_pick/internal/domain"
	"unnis_pick/internal/searchengine"
)

const QueryPrefix = "query"

type ProductService struct {
	queries       *query.Queries
	searchService searchengine.Service
	cacheService  cache.Service
}

type ProductDoc struct {
	Total int64 `json:"total"`
	Items any   `json:"items"`
}

func NewProductService(queries *query.Queries, searchService searchengine.Service, cacheService cache.Service) *ProductService {
	return &ProductService{
		queries:       queries,
		searchService: searchService,
		cacheService:  cacheService,
	}
}

func (s *ProductService) CreateProduct(ctx context.Context, product *domain.Product) (*domain.Product, error) {
	brandId, err := mapper.StringToUUID(product.BrandID)
	if err != nil {
		return nil, err
	}

	price, err := mapper.Float64ToNumeric(product.Price)
	if err != nil {
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

	err = s.cacheService.BatchDelByPrefix(ctx, QueryPrefix)
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

	err = s.cacheService.BatchDelByPrefix(ctx, QueryPrefix)
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

func (s *ProductService) QueryProducts(ctx context.Context, filter *domain.ProductFilter) (*domain.ProductQueryResponse, error) {
	key := buildQueryKey(filter)
	// if err := s.cacheService.Get(ctx, key, &res); err == nil {
	// 	return res, nil
	// }
	cache, err := s.retrieveQueryFromCache(ctx, key)
	if err == nil {
		return cache, nil
	}

	log.Println("Querying products from search service")

	queryResponse := s.searchService.QueryProducts(filter)
	resJson, err := json.Marshal(queryResponse)
	if err != nil {
		return nil, err
	}

	err = s.cacheService.Set(ctx, key, resJson, time.Hour)
	if err != nil {
		log.Printf("Failed to cache products: %v", err)
	}

	var res *domain.ProductQueryResponse
	err = json.Unmarshal(resJson, &res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func buildQueryKey(filter *domain.ProductFilter) string {

	key := fmt.Sprintf("%s:%s:%f:%f:%d:%d", filter.Brand, filter.Name, filter.PriceMin, filter.PriceMax, filter.PageNumber, filter.PageSize)

	hash := sha256.Sum256([]byte(key))
	return fmt.Sprintf("%s::%x", QueryPrefix, hash)
}

func (s *ProductService) retrieveQueryFromCache(ctx context.Context, key string) (*domain.ProductQueryResponse, error) {
	var res *domain.ProductQueryResponse

	cache, err := s.cacheService.Get(ctx, key)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal([]byte(cache), &res); err != nil {
		return nil, err
	}

	return res, nil

}
