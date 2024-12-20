package categories

import (
	entity "milestone2/entities"

	"github.com/stretchr/testify/mock"
)

type MockCategoryRepository struct {
	Mock mock.Mock
}

func (mockCategoryRepository *MockCategoryRepository) Find() ([]entity.Category, error) {
	args := mockCategoryRepository.Mock.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]entity.Category), args.Error(1)
}
