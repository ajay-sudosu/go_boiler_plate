package routes

import (
	"abc/internal/handler"

	"github.com/labstack/echo/v4"
)

func RegisterUserRoutes(e *echo.Group, h *handler.NetworkHandler) {
	e.POST("/users", h.CreateNetwork)
	e.GET("/users/:id", h.GetNetworks)

}
