package server

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	_ "github.com/joho/godotenv/autoload"

	"unnis_pick/internal/database"
	"unnis_pick/internal/domain"
	"unnis_pick/internal/service"
)

type Server struct {
	port           int
	dbPool         database.Service
	brandService   domain.BrandService
	productService domain.ProductService
}

func NewServer() *http.Server {
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	dbPool := database.New()
	brandService := service.NewBrandService(dbPool.Queries())
	productService := service.NewProductService(dbPool.Queries())

	NewServer := &Server{
		port:           port,
		dbPool:         dbPool,
		brandService:   brandService,
		productService: productService,
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
