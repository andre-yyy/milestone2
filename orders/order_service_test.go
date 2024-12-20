package orders

import (
	"errors"
	"testing"

	entity "milestone2/entities"
	dto "milestone2/orders/dtos"

	"github.com/stretchr/testify/assert"
)

func TestGetAllOrdersFailed(t *testing.T) {
	mockOrderRepository := new(MockOrderRepository)
	mockOrderRepository.Mock.On("Find").Return(nil, errors.New("failed to get all orders"))
	orderService := NewOrderService(mockOrderRepository)
	results, err := orderService.GetAllOrders()

	assert.Nil(t, results)
	assert.NotNil(t, err)
	assert.Equal(t, "failed to get all orders", err.Error(), "The slice should be nil when no orders are found")
}

func TestGetAllOrdersSuccess(t *testing.T) {
	orders := []entity.Order{
		{ID: 1, UserID: "44acc6ef-c8ad-4e91-9c15-3bfd1a27adce", ProductID: 1, RentDays: 7, TotalPrice: 27.5, Status: "pending"},
		{ID: 2, UserID: "ae54deb8-a68c-46eb-9459-35b4bc344392", ProductID: 2, RentDays: 15, TotalPrice: 52.5, Status: "delivered"},
	}
	mockOrderRepository := new(MockOrderRepository)
	mockOrderRepository.Mock.On("Find").Return(orders, nil)
	orderService := NewOrderService(mockOrderRepository)
	results, err := orderService.GetAllOrders()

	assert.Nil(t, err)
	assert.NotNil(t, results)
	assert.Equal(t, 1, orders[0].ProductID, "The first orders's id should be 1")
	assert.Equal(t, "delivered", orders[1].Status, "The second orders's status should be delivered")
}
func TestGetOrderByIdNotFound(t *testing.T) {
	mockOrderRepository := new(MockOrderRepository)
	mockOrderRepository.Mock.On("FindOneById", "1").Return(nil, errors.New("order with id '1' not found"))
	orderService := NewOrderService(mockOrderRepository)
	result, err := orderService.GetOrderById("1")

	assert.Nil(t, result)
	assert.NotNil(t, err)
	assert.Equal(t, "order with id '1' not found", err.Error(), "error message should be order with id '1' not found")
}
func TestGetOrderByIdSuccess(t *testing.T) {
	order := entity.Order{
		ID:        1,
		UserID:    "44acc6ef-c8ad-4e91-9c15-3bfd1a27adce",
		ProductID: 1,
		Status:    "pending",
	}
	mockOrderRepository := new(MockOrderRepository)
	mockOrderRepository.Mock.On("FindOneById", "1").Return(&order, nil)
	orderService := NewOrderService(mockOrderRepository)
	result, err := orderService.GetOrderById("1")

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, 1, result.ProductID, "result has to be order with product id 1")
	assert.Equal(t, "pending", result.Status, "result has to be order with status 'pending'")
	assert.Equal(t, &order, result, "expected order pointers to be equal")
}

func TestCreateOrderValidationError(t *testing.T) {
	orderToCreate := dto.CreateOrderDto{
		ProductID: 0,
		RentDays:  5,
	}
	newOrder := entity.Order{
		ProductID: orderToCreate.ProductID,
		RentDays:  orderToCreate.RentDays,
	}
	mockOrderRepository := new(MockOrderRepository)
	mockOrderRepository.Mock.On("Create", newOrder).Return(nil, errors.New("product id is required"))
	orderService := NewOrderService(mockOrderRepository)
	result, err := orderService.CreateOrder(newOrder)

	assert.Nil(t, result)
	assert.NotNil(t, err)
	assert.Equal(t, "product id is required", err.Error(), "error message should be 'product id is required'")
}

func TestCreateOrderSuccess(t *testing.T) {
	orderToCreate := dto.CreateOrderDto{
		ProductID: 1,
		RentDays:  5,
	}
	newOrder := entity.Order{
		ProductID: orderToCreate.ProductID,
		RentDays:  orderToCreate.RentDays,
	}
	mockOrderRepository := new(MockOrderRepository)
	mockOrderRepository.Mock.On("Create", newOrder).Return(&newOrder, nil)
	orderService := NewOrderService(mockOrderRepository)
	result, err := orderService.CreateOrder(newOrder)

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, orderToCreate.ProductID, result.ProductID, "created order should have the same product id")
	assert.Equal(t, orderToCreate.RentDays, result.RentDays, "created order should have the same rent days")
}

func TestUpdateOrderValidationError(t *testing.T) {
	orderToUpdate := dto.UpdateOrderStatusDto{
		Status: "delivered",
	}
	newOrder := entity.Order{
		Status: orderToUpdate.Status,
	}
	mockOrderRepository := new(MockOrderRepository)
	mockOrderRepository.Mock.On("UpdateStatusById", newOrder).Return(nil, errors.New("status is required"))
	orderService := NewOrderService(mockOrderRepository)
	result, err := orderService.UpdateOrderStatusById(newOrder)

	assert.Nil(t, result)
	assert.NotNil(t, err)
	assert.Equal(t, "status is required", err.Error(), "error message should be 'status is required'")
}
func TestUpdateOrderSuccess(t *testing.T) {
	orderToUpdate := dto.UpdateOrderStatusDto{
		Status: "delivered",
	}
	newOrder := entity.Order{
		Status: orderToUpdate.Status,
	}
	mockOrderRepository := new(MockOrderRepository)
	mockOrderRepository.Mock.On("UpdateStatusById", newOrder).Return(&newOrder, nil)
	orderService := NewOrderService(mockOrderRepository)
	result, err := orderService.UpdateOrderStatusById(newOrder)

	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, orderToUpdate.Status, result.Status, "updated order should have the same status")
}
