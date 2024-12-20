package coins

import (
	"fmt"
	"strconv"

	entity "milestone2/entities"
	repository "milestone2/repositories"

	"gorm.io/gorm"
)

type coinRepositoryImpl struct {
	DB *gorm.DB
}

func NewCoinRepository(db *gorm.DB) repository.CoinRepository {
	return &coinRepositoryImpl{DB: db}
}

func (coinRepository *coinRepositoryImpl) Find() ([]entity.Coin, error) {
	var coin []entity.Coin
	err := coinRepository.DB.Find(&coin).Error
	if err != nil {
		return nil, fmt.Errorf("failed to get all coins")
	}
	return coin, nil
}

func (coinRepository *coinRepositoryImpl) FindOneById(coinId string) (*entity.Coin, error) {
	var coin entity.Coin
	coinID, err := strconv.Atoi(coinId)
	if err != nil {
		return nil, err
	}
	err = coinRepository.DB.Where(&entity.Coin{ID: coinID}).Take(&coin).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound || coin.ID <= 0 {
			return nil, fmt.Errorf("coin with id '%s' not found", coinId)
		}
		return nil, err
	}
	return &coin, nil
}
