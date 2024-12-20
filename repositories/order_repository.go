package repository

import entity "milestone2/entities"

type OrderRepository interface {
	Find() ([]entity.Order, error)
	FindByUserId(userId string) ([]entity.Order, error)
	FindByStatus(status string, userId *string) ([]entity.Order, error)
	FindOverdue(userId *string) ([]entity.Order, error)
	FindOneById(orderId string) (*entity.Order, error)
	Create(Order entity.Order) (*entity.Order, error)
	UpdateStatusById(orderToUpdate entity.Order) (*entity.Order, error)
}
