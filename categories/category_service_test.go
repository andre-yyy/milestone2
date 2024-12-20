package categories

import (
	"errors"
	"testing"

	entity "milestone2/entities"

	"github.com/stretchr/testify/assert"
)

func TestGetAllCategoriesSuccess(t *testing.T) {
	categories := []entity.Category{
		{ID: 1, Name: "Wheelchairs"},
		{ID: 2, Name: "Walking Aids"},
	}
	mockCategoryRepository := new(MockCategoryRepository)
	mockCategoryRepository.Mock.On("Find").Return(categories, nil)
	categorySerice := NewCategoryService(mockCategoryRepository)
	results, err := categorySerice.GetAllCategories()

	assert.Nil(t, err)
	assert.NotNil(t, results)
	assert.Equal(t, "Wheelchairs", results[0].Name, "The first category's name should be Electric Wheelchair")
	assert.Equal(t, "Walking Aids", results[1].Name, "The first category's name should be Walking Aid")
}

func TestGetAllCategoriesFailed(t *testing.T) {
	mockCategoryRepository := new(MockCategoryRepository)
	mockCategoryRepository.Mock.On("Find").Return(nil, errors.New("failed to get all categories"))
	categoryService := NewCategoryService(mockCategoryRepository)
	result, err := categoryService.GetAllCategories()

	assert.Nil(t, result)
	assert.NotNil(t, err)
	assert.Equal(t, "failed to get all categories", err.Error(), "The slice should be nil when no categories are found")
}
