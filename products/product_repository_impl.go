package products

import (
	"fmt"
	"strconv"

	entity "milestone2/entities"
	repository "milestone2/repositories"

	"gorm.io/gorm"
)

type productRepositoryImpl struct {
	DB *gorm.DB
}

func NewProductRepository(db *gorm.DB) repository.ProductRepository {
	return &productRepositoryImpl{DB: db}
}

func (productRepository *productRepositoryImpl) Find() ([]entity.Product, error) {
	var product []entity.Product
	err := productRepository.DB.Find(&product).Error
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (productRepository *productRepositoryImpl) FindOneById(productId string) (*entity.Product, error) {
	var product entity.Product
	productID, err := strconv.Atoi(productId)
	if err != nil {
		return nil, err
	}
	err = productRepository.DB.Where(&entity.Product{ID: productID}).Take(&product).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound || product.ID <= 0 {
			return nil, fmt.Errorf("product with id '%s' not found", productId)
		}
		return nil, err
	}
	return &product, nil
}
