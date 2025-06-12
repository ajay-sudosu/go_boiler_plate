package routes

import (
	"abc/internal/di"

	"github.com/labstack/echo/v4"
)

func RegisterAllRoutes(api *echo.Group, c *di.Container) {
	// api := e.Group("/api/v1")

	// User routes
	RegisterUserRoutes(api, c.NetworkHandler)

	// Product routes
	RegisterProductRoutes(api, c.ProductHandler)
}
