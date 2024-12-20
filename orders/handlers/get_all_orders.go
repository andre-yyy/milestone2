package handlers

import (
	"net/http"

	entity "milestone2/entities"
	"milestone2/orders"
	"milestone2/users"

	"github.com/labstack/echo/v4"
)

// @summary Get all orders
// @description Get all orders
// @tags Milestone 2
// @security Auth
// @produce json
// @success 200 {object} entity.ServerResponseData[[]entity.Order]
// @failure 401 {object} entity.ServerResponse
// @failure 404 {object} entity.ServerResponse
// @failure 500 {object} entity.ServerResponse
// @router /orders [get]
func GetAllOrders(c echo.Context, os orders.OrderService, us users.UserService) error {
	userId := c.Get("user_id").(string)
	user, err := us.GetUserById(userId)
	if err != nil {
		response := entity.ServerResponse{
			Message: err.Error(),
			Code:    http.StatusNotFound,
		}
		return c.JSON(http.StatusNotFound, response)
	}

	if user.Role == "admin" {
		orders, err := os.GetAllOrders()
		if err != nil {
			response := entity.ServerResponse{
				Message: err.Error(),
				Code:    http.StatusInternalServerError,
			}
			return c.JSON(http.StatusInternalServerError, response)
		}
		response := entity.ServerResponseData[[]entity.Order]{
			Message: "success get all orders.",
			Code:    http.StatusOK,
			Data:    orders,
		}
		return c.JSON(http.StatusOK, response)
	}
	orders, err := os.GetAllOrdersByUserId(userId)
	if err != nil {
		response := entity.ServerResponse{
			Message: err.Error(),
			Code:    http.StatusInternalServerError,
		}
		return c.JSON(http.StatusInternalServerError, response)
	}
	response := entity.ServerResponseData[[]entity.Order]{
		Message: "success get all orders",
		Code:    http.StatusOK,
		Data:    orders,
	}
	return c.JSON(http.StatusOK, response)
}
