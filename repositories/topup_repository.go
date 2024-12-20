package repository

import entity "milestone2/entities"

type TopupRepository interface {
	Find() ([]entity.Topup, error)
	FindByUserId(userId string) ([]entity.Topup, error)
	FindByStatus(status string, userId *string) ([]entity.Topup, error)
	FindOneByExternalId(externalId string) (entity.Topup, error)
	Create(newTopup entity.Topup) (entity.Topup, error)
	UpdateStatusByExternalId(topupToUpdate entity.Topup) (entity.Topup, error)
	// DeleteById(topupId string) (entity.Topup, error)
}
