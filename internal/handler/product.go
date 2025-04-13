package handler

import (
	"net/http"

	"abc/internal/model"
	"abc/internal/usecase"

	"github.com/labstack/echo/v4"
)

type ProductHandler struct {
	usecase *usecase.ProductUsecase
}

func NewProductHandler(u *usecase.ProductUsecase) *ProductHandler {
	return &ProductHandler{usecase: u}
}

// CreateUser godoc
// @Summary      Create a user.
// @Description  Adds a new product to the system.
// @Description  ### Product Details
// @Description  This endpoint allows you to add a new product to the system.
// @Description  - **Name**: Name of the product
// @Description  - **Price**: Product price in USD
// @Description  - *Optional*: `description`, `tags`
// @Description  Ensure that:
// @Description  - The `product_id` is unique
// @Description  - Price must be a positive number
// @Tags         Products
// @Accept       json
// @Produce      json
// @Param        user  body      model.Product  true  "Product Payload"
// @Success      201   {object}  model.Product
// @Failure      400   {object}  map[string]interface{}
// @Failure      500   {object}  map[string]interface{}
// @Router       /product [post]
func (h *ProductHandler) CreateProduct(c echo.Context) error {
	var product model.Product
	if err := c.Bind(&product); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid body"})
	}
	err := h.usecase.CreateProduct(&product)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	return c.JSON(http.StatusCreated, product)
}

func (h *ProductHandler) GetProducts(c echo.Context) error {
	products, err := h.usecase.GetProducts()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, products)
}
