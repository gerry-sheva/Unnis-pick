package searchengine

import "unnis_pick/internal/domain"

func productToDoc(product *domain.Product, brand string) *Product {
	return &Product{
		ProductID:   product.ProductID,
		Brand:       brand,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Stock:       product.Stock,
	}
}
