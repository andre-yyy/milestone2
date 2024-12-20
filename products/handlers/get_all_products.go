package handlers

import (
	"net/http"

	entity "milestone2/entities"
	"milestone2/products"

	"github.com/labstack/echo/v4"
)

// @summary Get all products
// @description Get all products
// @tags Milestone 2
// @produce json
// @success 200 {object} entity.ServerResponseData[[]entity.Product]
// @failure 500 {object} entity.ServerResponse
// @router /products [get]
func GetAllProducts(c echo.Context, ps *products.ProductService) error {
	allProducts, err := ps.GetAllProducts()
	if err != nil {
		response := entity.ServerResponse{
			Message: err.Error(),
			Code:    http.StatusInternalServerError,
		}
		return c.JSON(http.StatusInternalServerError, response)
	}
	response := entity.ServerResponseData[[]entity.Product]{
		Message: "success get all products",
		Code:    http.StatusOK,
		Data:    allProducts,
	}
	return c.JSON(http.StatusOK, response)
}
