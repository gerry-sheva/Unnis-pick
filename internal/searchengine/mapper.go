package searchengine

import "unnis_pick/internal/domain"

func productToDoc(product *domain.Product, brand string) *domain.ProductQuery {
	return &domain.ProductQuery{
		ProductID:   product.ProductID,
		Brand:       brand,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Stock:       product.Stock,
	}
}
