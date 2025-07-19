package handler

import (
	"net/http"
	"unnis_pick/internal/domain"

	"github.com/labstack/echo/v4"
)

type ProductHandler struct {
	productService domain.ProductService
}

func NewProductHandler(productService domain.ProductService) *ProductHandler {
	return &ProductHandler{
		productService: productService,
	}
}

func (h *ProductHandler) CreateProduct(c echo.Context) error {
	product := new(domain.Product)
	if err := c.Bind(product); err != nil {
		return err
	}
	ctx := c.Request().Context()
	res, err := h.productService.CreateProduct(ctx, product)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, res)
}

func (h *ProductHandler) GetProduct(c echo.Context) error {
	id := c.Param("id")
	ctx := c.Request().Context()
	product, err := h.productService.GetProduct(ctx, id)

	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, product)
}

func (h *ProductHandler) UpdateProduct(c echo.Context) error {
	id := c.Param("id")
	product := new(domain.Product)
	if err := c.Bind(product); err != nil {
		return err
	}
	ctx := c.Request().Context()
	res, err := h.productService.UpdateProduct(ctx, id, product)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, res)
}

func (h *ProductHandler) DeleteProduct(c echo.Context) error {
	id := c.Param("id")
	ctx := c.Request().Context()
	err := h.productService.DeleteProduct(ctx, id)
	if err != nil {
		return err
	}
	return c.NoContent(http.StatusNoContent)
}

func (h *ProductHandler) QueryProducts(c echo.Context) error {
	filter := new(domain.ProductFilter)
	if err := c.Bind(filter); err != nil {
		return err
	}
	filter.SetDefault()
	err := filter.Validate()
	if err != nil {
		return err
	}
	ctx := c.Request().Context()
	res, err := h.productService.QueryProducts(ctx, filter)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, res)
}
