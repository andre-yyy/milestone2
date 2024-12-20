package coins

import (
	entity "milestone2/entities"

	"github.com/stretchr/testify/mock"
)

type MockCoinRepository struct {
	Mock mock.Mock
}

func (coinRepository *MockCoinRepository) Find() ([]entity.Coin, error) {
	args := coinRepository.Mock.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]entity.Coin), args.Error(1)
}

func (coinRepository *MockCoinRepository) FindOneById(coinId string) (*entity.Coin, error) {
	args := coinRepository.Mock.Called(coinId)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.Coin), args.Error(1)
}
