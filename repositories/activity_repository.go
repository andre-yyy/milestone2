package repository

import entity "milestone2/entities"

type ActivityRepository interface {
	Find() ([]entity.Activity, error)
	FindByUserId(userId string) ([]entity.Activity, error)
}
