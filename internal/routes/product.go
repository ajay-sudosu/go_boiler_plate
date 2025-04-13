package routes

import (
	"abc/internal/handler"

	"github.com/labstack/echo/v4"
)

func RegisterProductRoutes(e *echo.Group, h *handler.ProductHandler) {
	e.GET("/products", h.GetProducts)
	e.GET("/products/:id", h.CreateProduct)
}
