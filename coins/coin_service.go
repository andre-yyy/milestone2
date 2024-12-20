package coins

import (
	entity "milestone2/entities"
	repository "milestone2/repositories"
)

type CoinService struct {
	coinRepository repository.CoinRepository
}

func NewCoinService(coinRepository repository.CoinRepository) *CoinService {
	return &CoinService{coinRepository: coinRepository}
}

func (s *CoinService) GetAllCoins() ([]entity.Coin, error) {
	allCoins, err := s.coinRepository.Find()
	if err != nil {
		return nil, err
	}
	return allCoins, nil
}

func (s *CoinService) GetCoinById(coinId string) (entity.Coin, error) {
	coin, err := s.coinRepository.FindOneById(coinId)
	if err != nil {
		return entity.Coin{}, err
	}
	return *coin, nil
}
