package orders

import (
	entity "milestone2/entities"
	repository "milestone2/repositories"
)

type OrderService struct {
	orderRepository repository.OrderRepository
}

func NewOrderService(orderRepository repository.OrderRepository) *OrderService {
	return &OrderService{orderRepository: orderRepository}
}

func (s *OrderService) GetAllOrders() ([]entity.Order, error) {
	allOrders, err := s.orderRepository.Find()
	if err != nil {
		return nil, err
	}
	return allOrders, nil
}

func (s *OrderService) GetAllOrdersByUserId(userId string) ([]entity.Order, error) {
	allOrders, err := s.orderRepository.FindByUserId(userId)
	if err != nil {
		return nil, err
	}
	return allOrders, nil
}

func (s *OrderService) GetAllOrdersByStatus(status string, userId *string) ([]entity.Order, error) {
	allOrders, err := s.orderRepository.FindByUserId(*userId)
	if err != nil {
		return nil, err
	}
	return allOrders, nil
}

func (s *OrderService) GetOrderById(orderId string) (*entity.Order, error) {
	order, err := s.orderRepository.FindOneById(orderId)
	if err != nil {
		return nil, err
	}
	return order, nil
}

func (s *OrderService) CreateOrder(newOrder entity.Order) (*entity.Order, error) {
	createdOrder, err := s.orderRepository.Create(newOrder)
	if err != nil {
		return nil, err
	}
	return createdOrder, nil
}

func (s *OrderService) UpdateOrderStatusById(orderToUpdate entity.Order) (*entity.Order, error) {
	updatedOrder, err := s.orderRepository.UpdateStatusById(orderToUpdate)
	if err != nil {
		return nil, err
	}
	return updatedOrder, nil
}

func (s *OrderService) GetAllOverdueOrders(userId *string) ([]entity.Order, error) {
	allOrders, err := s.orderRepository.FindOverdue(userId)
	if err != nil {
		return nil, err
	}
	return allOrders, nil
}
