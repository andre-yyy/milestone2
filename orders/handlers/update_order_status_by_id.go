package handlers

import (
	"fmt"
	"net/http"

	entity "milestone2/entities"
	"milestone2/orders"
	dto "milestone2/orders/dtos"
	"milestone2/users"

	"github.com/labstack/echo/v4"
)

// @summary Update order status by id
// @description Update order status by id
// @tags Milestone 2
// @security Auth
// @accept json
// @produce json
// @param orderID path int true "orderID"
// @param order body dto.UpdateOrderStatusDto true "Order"
// @success 200 {object} entity.ServerResponseData[entity.Order]
// @failure 400 {object} entity.ServerResponse
// @failure 401 {object} entity.ServerResponse
// @failure 403 {object} entity.ServerResponse
// @failure 404 {object} entity.ServerResponse
// @failure 500 {object} entity.ServerResponse
// @router /orders/{orderID} [patch]
func UpdateOrderStatusById(c echo.Context, os *orders.OrderService, us users.UserService) error {
	orderId := c.Param("id")
	userId := c.Get("user_id").(string)
	var newOrderStatus dto.UpdateOrderStatusDto
	err := c.Bind(&newOrderStatus)
	if err != nil {
		response := entity.ServerResponse{
			Message: "invalid json",
			Code:    http.StatusBadRequest,
		}
		return c.JSON(http.StatusBadRequest, response)
	}

	order, err := os.GetOrderById(orderId)
	if err != nil {
		response := entity.ServerResponse{
			Message: err.Error(),
			Code:    http.StatusNotFound,
		}
		return c.JSON(http.StatusNotFound, response)
	}

	if newOrderStatus.Status == "" {
		response := entity.ServerResponse{
			Message: "status is required",
			Code:    http.StatusBadRequest,
		}
		return c.JSON(http.StatusBadRequest, response)
	}

	user, err := us.GetUserById(userId)
	if err != nil {
		response := entity.ServerResponse{
			Message: err.Error(),
			Code:    http.StatusNotFound,
		}
		return c.JSON(http.StatusNotFound, response)
	}

	fmt.Println("User ID :", user.ID)
	fmt.Println("user ID order :", order.UserID)

	if user.ID != order.UserID {
		response := entity.ServerResponse{
			Message: "cannot modify order that created by another user",
			Code:    http.StatusForbidden,
		}
		return c.JSON(http.StatusForbidden, response)
	}

	orderToUpdate := entity.Order{
		ID:     order.ID,
		Status: newOrderStatus.Status,
	}

	updatedOrder, err := os.UpdateOrderStatusById(orderToUpdate)
	if err != nil {
		response := entity.ServerResponse{
			Message: err.Error(),
			Code:    http.StatusInternalServerError,
		}
		return c.JSON(http.StatusInternalServerError, response)
	}

	response := entity.ServerResponseData[entity.Order]{
		Message: "success update order status",
		Code:    http.StatusOK,
		Data:    *updatedOrder,
	}
	return c.JSON(http.StatusOK, response)
}
