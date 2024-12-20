package handlers

import (
	"fmt"
	"math"
	"net/http"
	"strconv"

	entity "milestone2/entities"
	"milestone2/orders"
	dto "milestone2/orders/dtos"
	"milestone2/products"
	"milestone2/users"

	"github.com/labstack/echo/v4"
)

// @summary Create order
// @description Create order
// @tags Milestone 2
// @security Auth
// @accept json
// @produce json
// @param order body dto.CreateOrderDto true "Order"
// @success 201 {object} entity.ServerResponseData[entity.Order]
// @failure 400 {object} entity.ServerResponse
// @failure 401 {object} entity.ServerResponse
// @failure 404 {object} entity.ServerResponse
// @failure 500 {object} entity.ServerResponse
// @router /orders [post]
func CreateOrder(c echo.Context, os orders.OrderService, ps products.ProductService, us users.UserService) error {
	var orderToCreate dto.CreateOrderDto
	userId := c.Get("user_id").(string)
	user, err := us.GetUserById(userId)
	if err != nil {
		response := entity.ServerResponse{
			Message: err.Error(),
			Code:    http.StatusNotFound,
		}
		return c.JSON(http.StatusNotFound, response)
	}

	err = c.Bind(&orderToCreate)
	if err != nil {
		response := entity.ServerResponse{
			Message: "invalid json",
			Code:    http.StatusBadRequest,
		}
		return c.JSON(http.StatusBadRequest, response)
	}

	if orderToCreate.ProductID <= 0 {
		response := entity.ServerResponse{
			Message: "product id is required",
			Code:    http.StatusBadRequest,
		}
		return c.JSON(http.StatusBadRequest, response)
	}

	if orderToCreate.RentDays <= 0 {
		response := entity.ServerResponse{
			Message: "rent days must be higher than 0",
			Code:    http.StatusBadRequest,
		}
		return c.JSON(http.StatusBadRequest, response)
	}

	productID := strconv.Itoa(orderToCreate.ProductID)
	product, err := ps.GetProductById(productID)
	if err != nil {
		response := entity.ServerResponse{
			Message: err.Error(),
			Code:    http.StatusNotFound,
		}
		return c.JSON(http.StatusNotFound, response)
	}

	if product.Stock <= 0 {
		response := entity.ServerResponse{
			Message: "out of stock",
			Code:    http.StatusBadRequest,
		}
		return c.JSON(http.StatusBadRequest, response)
	}

	totalPrice := product.RentalCost + (float64(orderToCreate.RentDays) * product.Daily)
	coinsToAdd := math.Abs(totalPrice - user.Deposit)

	if user.Deposit <= totalPrice {
		response := entity.ServerResponse{
			Message: fmt.Sprintf("Please add more coins. Currently, you have %v coins in your account. You will need to add %v coins or more to complete this order, with base rental cost %v coins and daily fees %v coins/day for %v rental day(s)", user.Deposit, coinsToAdd, product.RentalCost, product.Daily, orderToCreate.RentDays),
			Code:    http.StatusBadRequest,
		}
		return c.JSON(http.StatusBadRequest, response)
	}

	if user.Address == "" || user.CityID == "" {
		response := entity.ServerResponse{
			Message: "please provide your location information",
			Code:    http.StatusBadRequest,
		}
		return c.JSON(http.StatusBadRequest, response)
	}

	newOrder := entity.Order{
		UserID:        user.ID,
		ProductID:     orderToCreate.ProductID,
		RentDays:      orderToCreate.RentDays,
		TotalPrice:    totalPrice,
		DestinationID: user.CityID,
	}

	createdOrder, err := os.CreateOrder(newOrder)
	if err != nil {
		response := entity.ServerResponse{
			Message: err.Error(),
			Code:    http.StatusInternalServerError,
		}
		return c.JSON(http.StatusInternalServerError, response)
	}

	response := entity.ServerResponseData[entity.Order]{
		Message: "success create order",
		Code:    http.StatusCreated,
		Data:    *createdOrder,
	}
	return c.JSON(http.StatusCreated, response)
}
