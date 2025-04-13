package routes

import (
	"abc/internal/handler"

	"github.com/labstack/echo/v4"
)

func RegisterUserRoutes(e *echo.Group, h *handler.UserHandler) {
	e.GET("/users", h.CreateUser)
	e.GET("/users/:id", h.GetUsers)
}
