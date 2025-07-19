package searchengine

import (
	"context"
	"log"
	"os"
	"unnis_pick/internal/domain"

	"github.com/elastic/go-elasticsearch/v9"
	"github.com/elastic/go-elasticsearch/v9/typedapi/core/search"
	"github.com/elastic/go-elasticsearch/v9/typedapi/indices/create"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Service interface {
	SetupProductsIndex()
	QueryProducts(filter *domain.ProductFilter) *search.Response
	IndexProduct(product *domain.Product, brand string) error
	DeleteProduct(id string)
}

type service struct {
	search *elasticsearch.TypedClient
}

var (
	cloudId        = os.Getenv("SEARCH_CLOUD_ID")
	apiKey         = os.Getenv("SEARCH_API_KEY")
	searchInstance *service
)

func New() Service {
	// Reuse client
	if searchInstance != nil {
		return searchInstance
	}

	typedClient, err := elasticsearch.NewTypedClient(elasticsearch.Config{
		Addresses: []string{cloudId},
		APIKey:    apiKey,
	})
	if err != nil {
		log.Fatal("Failed to create Elasticsearch client:", err)
	}

	searchInstance = &service{
		search: typedClient,
	}

	return searchInstance
}

func (s *service) SetupProductsIndex() {
	exists, err := s.search.Indices.Exists("products").Do(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	if exists {
		return
	}

	_, err = s.search.Indices.Create("products").
		Request(&create.Request{
			Mappings: &types.TypeMapping{
				Properties: map[string]types.Property{
					"productId": types.NewTextProperty(),
					"name":      types.NewTextProperty(),
					"price":     types.NewScaledFloatNumberProperty(),
					"stock":     types.NewIntegerNumberProperty(),
					"brand":     types.NewKeywordProperty(),
				},
			},
		}).Do(context.Background())
	if err != nil {
		log.Fatal(err)
	}
}

func (s *service) QueryProducts(filter *domain.ProductFilter) *search.Response {
	mustQueries := []types.Query{}

	if filter.Name != "" {
		mustQueries = append(mustQueries, types.Query{
			Match: map[string]types.MatchQuery{
				"name": {Query: filter.Name},
			},
		})
	}

	if filter.Brand != "" {
		mustQueries = append(mustQueries, types.Query{
			Match: map[string]types.MatchQuery{
				"brand": {Query: filter.Brand},
			},
		})
	}

	if filter.PriceMin > 0 || filter.PriceMax > 0 {
		rangeQuery := types.NumberRangeQuery{}
		if filter.PriceMin > 0 {
			min := types.Float64(filter.PriceMin)
			rangeQuery.Gte = &min
		}
		if filter.PriceMax > 0 {
			log.Println(filter.PriceMax)
			max := types.Float64(filter.PriceMax)
			rangeQuery.Lte = &max
		}

		mustQueries = append(mustQueries, types.Query{
			Range: map[string]types.RangeQuery{
				"price": rangeQuery,
			},
		})
	}

	query := &types.Query{
		Bool: &types.BoolQuery{
			Must: mustQueries,
		},
	}

	from := filter.PageNumber * filter.PageSize
	size := filter.PageSize
	res, err := s.search.Search().
		Index("products").
		Request(&search.Request{
			From:  &from,
			Size:  &size,
			Query: query,
		}).Do(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	return res
}

func (s *service) IndexProduct(product *domain.Product, brand string) error {
	doc := productToDoc(product, brand)
	_, err := s.search.Index("products").
		Id(product.ProductID).
		Request(doc).
		Do(context.Background())
	if err != nil {
		return err
	}

	return nil
}

func (s *service) DeleteProduct(id string) {
	s.search.Delete("products", id).Do(context.Background())
}
