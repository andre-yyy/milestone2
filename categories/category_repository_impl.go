package categories

import (
	"fmt"

	entity "milestone2/entities"
	repository "milestone2/repositories"

	"gorm.io/gorm"
)

type categoryRepositoryImpl struct {
	DB *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) repository.CategoryRepository {
	return &categoryRepositoryImpl{DB: db}
}

func (categoryRepository *categoryRepositoryImpl) Find() ([]entity.Category, error) {
	var category []entity.Category
	err := categoryRepository.DB.Find(&category).Error
	if err != nil {
		return nil, fmt.Errorf("failed to get all categories")
	}
	return category, nil
}
