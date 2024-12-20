package repository

import entity "milestone2/entities"

type CoinRepository interface {
	Find() ([]entity.Coin, error)
	FindOneById(coinId string) (*entity.Coin, error)
}
