package coins

import (
	"errors"
	"testing"

	entity "milestone2/entities"

	"github.com/stretchr/testify/assert"
)

func TestGetAllCoinsFailed(t *testing.T) {
	mockCoinRepository := new(MockCoinRepository)
	mockCoinRepository.Mock.On("Find").Return(nil, errors.New("failed to get all coins"))
	coinService := NewCoinService(mockCoinRepository)
	results, err := coinService.GetAllCoins()

	assert.Nil(t, results)
	assert.NotNil(t, err)
	assert.Equal(t, "failed to get all coins", err.Error(), "The slice should be nil when no coins are found")
}
func TestGetAllCoinsSuccess(t *testing.T) {
	coins := []entity.Coin{
		{ID: 1, Name: "10 coins", Description: "increase your deposit by 10", Price: 10.00},
		{ID: 2, Name: "25 coins", Description: "increase your deposit by 25", Price: 25.00},
	}
	mockCoinRepository := new(MockCoinRepository)
	mockCoinRepository.Mock.On("Find").Return(coins, nil)
	coinService := NewCoinService(mockCoinRepository)
	results, err := coinService.GetAllCoins()

	assert.Nil(t, err)
	assert.NotNil(t, results)
	assert.Equal(t, "10 coins", results[0].Name, "The first coins' name should be 10 coins")
	assert.Equal(t, 25.00, results[1].Price, "The second coins' price should be 25")
}
