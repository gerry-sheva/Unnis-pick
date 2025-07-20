package server

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	_ "github.com/joho/godotenv/autoload"

	"unnis_pick/internal/cache"
	"unnis_pick/internal/database"
	"unnis_pick/internal/domain"
	"unnis_pick/internal/searchengine"
	"unnis_pick/internal/service"
)

type Server struct {
	port           int
	dbPool         database.Service
	brandService   domain.BrandService
	productService domain.ProductService
	searchService  searchengine.Service
}

func NewServer() *http.Server {
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	dbPool := database.New()
	searchService := searchengine.New()
	searchService.SetupProductsIndex()
	cacheService := cache.New()

	brandService := service.NewBrandService(dbPool.Queries())
	productService := service.NewProductService(dbPool.Queries(), searchService, cacheService)

	NewServer := &Server{
		port:           port,
		dbPool:         dbPool,
		brandService:   brandService,
		productService: productService,
		searchService:  searchService,
	}

	// Declare Server config
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", NewServer.port),
		Handler:      NewServer.RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server
}
