package handlers

import (
	"net/http"
	"strings"

	entity "milestone2/entities"
	"milestone2/products"

	"github.com/labstack/echo/v4"
)

// @summary Get product by id
// @description Get product by id
// @tags Milestone 2
// @produce json
// @param productID path int true "productID"
// @success 200 {object} entity.TopupInvoice[entity.Product]
// @failure 404 {object} entity.ServerResponse
// @failure 500 {object} entity.ServerResponse
// @router /products/{productID} [get]
func GetProductById(c echo.Context, ps *products.ProductService) error {
	productId := c.Param("id")
	product, err := ps.GetProductById(productId)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			response := entity.ServerResponse{
				Message: err.Error(),
				Code:    http.StatusNotFound,
			}
			return c.JSON(http.StatusNotFound, response)
		} else {
			response := entity.ServerResponse{
				Message: err.Error(),
				Code:    http.StatusInternalServerError,
			}
			return c.JSON(http.StatusInternalServerError, response)
		}
	}
	response := entity.ServerResponseData[entity.Product]{
		Message: "success get product",
		Code:    http.StatusOK,
		Data:    *product,
	}
	return c.JSON(http.StatusOK, response)
}
