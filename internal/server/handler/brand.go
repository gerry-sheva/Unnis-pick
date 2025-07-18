package handler

import (
	"net/http"
	"unnis_pick/internal/domain"

	"github.com/labstack/echo/v4"
)

type BrandHandler struct {
	brandService domain.BrandService
}

func NewBrandHandler(brandService domain.BrandService) *BrandHandler {
	return &BrandHandler{
		brandService: brandService,
	}
}

func (h *BrandHandler) CreateBrand(c echo.Context) error {
	ctx := c.Request().Context()

	req := new(domain.Brand)
	if err := c.Bind(req); err != nil {
		return err
	}

	res, err := h.brandService.CreateBrand(ctx, req)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, res)
}

func (h *BrandHandler) GetBrand(c echo.Context) error {
	id := c.Param("id")
	ctx := c.Request().Context()

	res, err := h.brandService.GetBrand(ctx, id)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, res)
}

func (h *BrandHandler) UpdateBrand(c echo.Context) error {
	id := c.Param("id")
	ctx := c.Request().Context()

	req := new(domain.Brand)
	if err := c.Bind(req); err != nil {
		return err
	}

	res, err := h.brandService.UpdateBrand(ctx, id, req)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, res)
}

func (h *BrandHandler) DeleteBrand(c echo.Context) error {
	id := c.Param("id")
	ctx := c.Request().Context()

	err := h.brandService.DeleteBrand(ctx, id)
	if err != nil {
		return err
	}
	return c.NoContent(http.StatusNoContent)
}
