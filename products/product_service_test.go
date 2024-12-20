package products

import (
	"errors"
	"testing"

	entity "milestone2/entities"

	"github.com/stretchr/testify/assert"
)

func TestGetAllProductsFailed(t *testing.T) {

	mockProductRepository := new(MockProductRepository)

	mockProductRepository.Mock.On("Find").Return(nil, errors.New("failed to get all products"))
	productService := NewProductService(mockProductRepository)
	results, err := productService.GetAllProducts()

	assert.Nil(t, results)
	assert.NotNil(t, err)
	assert.Equal(t, "failed to get all products", err.Error(), "The slice should be nil when no products are found")
}

func TestGetAllProductsSuccess(t *testing.T) {
	products := []entity.Product{
		{ID: 1, Name: "Wheelchair One", Stock: 10, RentalCost: 10},
		{ID: 2, Name: "Wheelchair Two", Stock: 15, RentalCost: 5},
	}

	mockProductRepository := new(MockProductRepository)
	mockProductRepository.Mock.On("Find").Return(products, nil)
	productService := NewProductService(mockProductRepository)
	results, err := productService.GetAllProducts()

	assert.NotNil(t, results)
	assert.Nil(t, err)
	assert.Equal(t, products[0].ID, results[0].ID, "The first product's name should be Electric Wheelchair")
	assert.Equal(t, products[1].Name, results[1].Name, "The second product's name should be Light Walking Aid")
	assert.Equal(t, products, results, "The second product's name should be Light Walking Aid")

}

func TestGetProductByIdNotFound(t *testing.T) {
	mockProductRepository := new(MockProductRepository)
	mockProductRepository.Mock.On("FindOneById", "100").Return(nil, errors.New("product with id '1' not found"))

	productService := NewProductService(mockProductRepository)
	result, err := productService.GetProductById("100")
	assert.Nil(t, result)
	assert.NotNil(t, err)
	assert.Equal(t, "product with id '1' not found", err.Error(), "Test error product not found")
}
func TestGetProductByIdSuccess(t *testing.T) {

	product := entity.Product{
		ID:         1,
		Name:       "Wheelchair",
		Stock:      5,
		RentalCost: 5,
		Daily:      2.5,
		CategoryID: 1,
	}

	mockProductRepository := new(MockProductRepository)
	mockProductRepository.Mock.On("FindOneById", "1").Return(&product, nil)

	productService := NewProductService(mockProductRepository)
	result, err := productService.GetProductById("1")
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, 1, result.ID)
	assert.Equal(t, "Wheelchair", result.Name)
}
