package orders

import (
	"fmt"
	"strconv"
	"time"

	entity "milestone2/entities"
	repository "milestone2/repositories"

	"gorm.io/gorm"
)

type orderRepositoryImpl struct {
	DB *gorm.DB
}

func NewOrderRepository(db *gorm.DB) repository.OrderRepository {
	return &orderRepositoryImpl{DB: db}
}

func (orderRepository *orderRepositoryImpl) Find() ([]entity.Order, error) {
	var orders []entity.Order
	err := orderRepository.DB.Find(&orders).Error
	if err != nil {
		return nil, fmt.Errorf("failed to get all orders")
	}
	return orders, nil
}

func (orderRepository *orderRepositoryImpl) FindOneById(orderId string) (*entity.Order, error) {
	var order entity.Order
	orderID, err := strconv.Atoi(orderId)
	if err != nil {
		return nil, err
	}
	err = orderRepository.DB.Where(&entity.Order{ID: orderID}).Take(&order).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound || order.ID <= 0 {
			return nil, fmt.Errorf("order with id '%s' not found", orderId)
		}
		return nil, err
	}
	return &order, nil
}

func (orderRepository *orderRepositoryImpl) FindByUserId(userId string) ([]entity.Order, error) {
	var orders []entity.Order
	err := orderRepository.DB.Where(&entity.Order{UserID: userId}).Find(&orders).Error
	if err != nil {
		return nil, err
	}
	return orders, nil
}

func (orderRepository *orderRepositoryImpl) FindByStatus(status string, userId *string) ([]entity.Order, error) {
	var orders []entity.Order
	if userId == nil {
		err := orderRepository.DB.Where(&entity.Order{Status: status}).Find(&orders).Error
		if err != nil {
			return nil, err
		}
	} else {
		err := orderRepository.DB.Where(&entity.Order{Status: status, UserID: *userId}).Find(&orders).Error
		if err != nil {
			return nil, err
		}
	}
	return orders, nil
}

func (orderRepository *orderRepositoryImpl) Create(newOrder entity.Order) (*entity.Order, error) {
	err := orderRepository.DB.Create(&newOrder).Error
	if err != nil {
		return nil, fmt.Errorf("failed to create order : '%s'", err)
	}
	return &newOrder, nil
}

func (orderRepository *orderRepositoryImpl) UpdateStatusById(orderToUpdate entity.Order) (*entity.Order, error) {
	orderID := strconv.Itoa(orderToUpdate.ID)
	order, err := orderRepository.FindOneById(orderID)
	if err != nil {
		return nil, err
	}

	order.Status = orderToUpdate.Status

	if orderToUpdate.Status == "delivered" {
		order.StartDate = time.Now()
		order.EndDate = time.Now().Add(time.Duration(order.RentDays) * 24 * time.Hour)
	}

	err = orderRepository.DB.Save(&order).Error
	if err != nil {
		return nil, err
	}

	return order, nil
}

func (orderRepository *orderRepositoryImpl) FindOverdue(userId *string) ([]entity.Order, error) {
	now := time.Now()
	var orders []entity.Order
	if userId == nil {
		err := orderRepository.DB.Where("end_date < $1 and status = $2", now, "delivered").Find(&orders).Error
		if err != nil {
			return nil, err
		}
	} else {
		err := orderRepository.DB.Where("end_date < $1 and status = $2 and user_id = $3", now, "delivered", userId).Find(&orders).Error
		if err != nil {
			return nil, err
		}
	}
	return orders, nil
}
