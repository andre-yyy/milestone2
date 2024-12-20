package activities

import (
	entity "milestone2/entities"
	repository "milestone2/repositories"

	"gorm.io/gorm"
)

type activityRepositoryImpl struct {
	DB *gorm.DB
}

func NewActivityRepository(db *gorm.DB) repository.ActivityRepository {
	return &activityRepositoryImpl{DB: db}
}

func (ActivityRepository *activityRepositoryImpl) Find() ([]entity.Activity, error) {
	var activities []entity.Activity
	err := ActivityRepository.DB.Find(&activities).Error
	if err != nil {
		return nil, err
	}
	return activities, nil
}

func (ActivityRepository *activityRepositoryImpl) FindByUserId(userId string) ([]entity.Activity, error) {
	var activities []entity.Activity
	err := ActivityRepository.DB.Where(&entity.Activity{UserID: userId}).Find(&activities).Error
	if err != nil {
		return nil, err
	}
	return activities, nil
}
