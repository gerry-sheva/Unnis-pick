package server

import (
	"net/http"
	"unnis_pick/internal/server/handler"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func (s *Server) RegisterRoutes() http.Handler {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"https://*", "http://*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowHeaders:     []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	brandHandler := handler.NewBrandHandler(s.brandService)
	productHandler := handler.NewProductHandler(s.productService)

	e.POST("/brands", brandHandler.CreateBrand)
	e.GET("/brands/:id", brandHandler.GetBrand)
	e.PUT("/brands/:id", brandHandler.UpdateBrand)
	e.DELETE("/brands/:id", brandHandler.DeleteBrand)

	e.POST("/products", productHandler.CreateProduct)
	e.GET("/products/:id", productHandler.GetProduct)
	e.PUT("/products/:id", productHandler.UpdateProduct)
	e.DELETE("/products/:id", productHandler.DeleteProduct)

	e.GET("/", s.HelloWorldHandler)
	e.GET("/health", s.healthHandler)

	return e
}

func (s *Server) HelloWorldHandler(c echo.Context) error {
	resp := map[string]string{
		"message": "Hello World",
	}

	return c.JSON(http.StatusOK, resp)
}

func (s *Server) healthHandler(c echo.Context) error {
	return c.JSON(http.StatusOK, s.dbPool.Health())
}
