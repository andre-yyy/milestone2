package products

import (
	entity "milestone2/entities"

	"github.com/stretchr/testify/mock"
)

type MockProductRepository struct {
	Mock mock.Mock
}

func (mockProductRepository *MockProductRepository) Find() ([]entity.Product, error) {
	args := mockProductRepository.Mock.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]entity.Product), args.Error(1)
}

func (mockProductRepository *MockProductRepository) FindOneById(productId string) (*entity.Product, error) {
	args := mockProductRepository.Mock.Called(productId)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.Product), args.Error(1)
}
