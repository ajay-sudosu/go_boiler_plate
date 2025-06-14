package handler

import (
	"net/http"

	"abc/internal/model"
	"abc/internal/usecase"

	"github.com/labstack/echo/v4"
)

type NetworkHandler struct {
	usecase *usecase.NetworkUsecase
}

func NewUserHandler(u *usecase.NetworkUsecase) *NetworkHandler {
	return &NetworkHandler{usecase: u}
}

// CreateUser godoc
// @Summary      Create a user
// @Description  Adds a new user to the system
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        user  body      model.User  true  "User Payload"
// @Success      201   {object}  model.User
// @Failure      400   {object}  map[string]interface{}
// @Failure      500   {object}  map[string]interface{}
// @Router       /users [post]
func (h *NetworkHandler) CreateNetwork(c echo.Context) error {

	var user model.User
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid body"})
	}
	err := h.usecase.CreateNetwork(&user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	return c.JSON(http.StatusCreated, user)
}

func (h *NetworkHandler) GetNetworks(c echo.Context) error {
	users, err := h.usecase.GetNetworks()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, users)
}
