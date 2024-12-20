package topups

import (
	"fmt"

	entity "milestone2/entities"
	repository "milestone2/repositories"

	"gorm.io/gorm"
)

type topupRepositoryImpl struct {
	DB *gorm.DB
}

func NewTopupRepository(db *gorm.DB) repository.TopupRepository {
	return &topupRepositoryImpl{DB: db}
}

func (topupRepository *topupRepositoryImpl) Find() ([]entity.Topup, error) {
	var topup []entity.Topup
	err := topupRepository.DB.Find(&topup).Error
	if err != nil {
		return nil, err
	}
	return topup, nil
}

func (topupRepository *topupRepositoryImpl) FindByUserId(userId string) ([]entity.Topup, error) {
	var topup []entity.Topup
	err := topupRepository.DB.Where(&entity.Topup{UserID: userId}).Find(&topup).Error
	if err != nil {
		return nil, err
	}
	return topup, nil
}

func (topupRepository *topupRepositoryImpl) FindByStatus(status string, userId *string) ([]entity.Topup, error) {
	var topup []entity.Topup
	if userId == nil {
		err := topupRepository.DB.Where(&entity.Topup{Status: status}).Find(&topup).Error
		if err != nil {
			return nil, err
		}
	} else {
		err := topupRepository.DB.Where(&entity.Topup{Status: status, UserID: *userId}).Find(&topup).Error
		if err != nil {
			return nil, err
		}
	}
	return topup, nil
}

func (topupRepository *topupRepositoryImpl) FindOneByExternalId(externalId string) (entity.Topup, error) {
	var topup entity.Topup
	err := topupRepository.DB.Where(&entity.Topup{ExternalID: externalId}).Take(&topup).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound || topup.ID <= 0 {
			return entity.Topup{}, fmt.Errorf("topup with external id '%s' not found", externalId)
		}
		return entity.Topup{}, err
	}
	return topup, nil
}

func (topupRepository *topupRepositoryImpl) Create(newTopup entity.Topup) (entity.Topup, error) {
	err := topupRepository.DB.Create(&newTopup).Error
	if err != nil {
		return entity.Topup{}, fmt.Errorf("failed to create topup : '%s'", err)
	}
	return newTopup, nil
}

func (topupRepository *topupRepositoryImpl) UpdateStatusByExternalId(topupToUpdate entity.Topup) (entity.Topup, error) {
	topup, err := topupRepository.FindOneByExternalId(topupToUpdate.ExternalID)
	if err != nil {
		return entity.Topup{}, err
	}
	topup.Status = topupToUpdate.Status
	err = topupRepository.DB.Save(&topup).Error
	if err != nil {
		return entity.Topup{}, err
	}
	return topup, nil
}
