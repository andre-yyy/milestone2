package categories

import (
	entity "milestone2/entities"
	repository "milestone2/repositories"
)

type CategoryService struct {
	categoryRepository repository.CategoryRepository
}

func NewCategoryService(categoryRepository repository.CategoryRepository) *CategoryService {
	return &CategoryService{categoryRepository: categoryRepository}
}

func (s *CategoryService) GetAllCategories() ([]entity.Category, error) {
	allCategories, err := s.categoryRepository.Find()
	if err != nil {
		return nil, err
	}
	return allCategories, nil
}
