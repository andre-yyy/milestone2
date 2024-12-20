package products

import (
	entity "milestone2/entities"
	repository "milestone2/repositories"
)

type ProductService struct {
	productRepository repository.ProductRepository
}

func NewProductService(productRepository repository.ProductRepository) *ProductService {
	return &ProductService{productRepository: productRepository}
}

func (s *ProductService) GetAllProducts() ([]entity.Product, error) {
	allProducts, err := s.productRepository.Find()
	if err != nil {
		return nil, err
	}
	return allProducts, nil
}

func (s *ProductService) GetProductById(productId string) (*entity.Product, error) {
	product, err := s.productRepository.FindOneById(productId)
	if err != nil {
		return nil, err
	}
	return product, nil
}
