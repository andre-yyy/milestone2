package repository

import entity "milestone2/entities"

type ProductRepository interface {
	Find() ([]entity.Product, error)
	FindOneById(productId string) (*entity.Product, error)
	// Create(newProduct entity.Product) (entity.Product, error)
	// UpdateById(productToUpdate entity.Product) (entity.Product, error)
	// DeleteById(productId string) (entity.Product, error)
}
