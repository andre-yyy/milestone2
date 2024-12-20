package orders

import (
	entity "milestone2/entities"

	"github.com/stretchr/testify/mock"
)

type MockOrderRepository struct {
	Mock mock.Mock
}

func (mockOrderRepository *MockOrderRepository) Find() ([]entity.Order, error) {
	args := mockOrderRepository.Mock.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]entity.Order), args.Error(1)
}

func (mockOrderRepository *MockOrderRepository) FindByUserId(userId string) ([]entity.Order, error) {
	args := mockOrderRepository.Mock.Called(userId)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]entity.Order), args.Error(1)
}

func (mockOrderRepository *MockOrderRepository) FindByStatus(status string, userId *string) ([]entity.Order, error) {
	args := mockOrderRepository.Mock.Called(userId)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]entity.Order), args.Error(1)
}

func (mockOrderRepository *MockOrderRepository) FindOneById(orderId string) (*entity.Order, error) {
	args := mockOrderRepository.Mock.Called(orderId)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.Order), args.Error(1)
}

func (mockOrderRepository *MockOrderRepository) Create(orderToCreate entity.Order) (*entity.Order, error) {
	args := mockOrderRepository.Mock.Called(orderToCreate)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.Order), args.Error(1)
}

func (mockOrderRepository *MockOrderRepository) UpdateStatusById(orderToUpdate entity.Order) (*entity.Order, error) {
	args := mockOrderRepository.Mock.Called(orderToUpdate)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.Order), args.Error(1)
}

func (mockOrderRepository *MockOrderRepository) FindOverdue(userId *string) ([]entity.Order, error) {
	args := mockOrderRepository.Mock.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]entity.Order), args.Error(1)
}
