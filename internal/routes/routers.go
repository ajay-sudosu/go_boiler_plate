package routes

import (
	"abc/internal/handler"

	"github.com/labstack/echo/v4"
)

func RegisterUserRoutes(g *echo.Group, h *handler.UserHandler) {
	g.POST("/users", h.CreateUser)
	g.GET("/users", h.GetUsers)
}

func RegisterProductRoutes(g *echo.Group, h *handler.ProductHandler) {
	g.POST("/products", h.CreateProduct)
	g.GET("/products", h.GetProducts)
}
