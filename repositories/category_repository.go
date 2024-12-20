package repository

import entity "milestone2/entities"

type CategoryRepository interface {
	Find() ([]entity.Category, error)
}
