package repository

import (
	entity "milestone2/entities"
	dto "milestone2/users/dtos"
)

type UserRepository interface {
	Find() ([]entity.User, error)
	FindOneById(userId string) (entity.User, error)
	FindOneByEmail(userEmail string) (entity.User, error)
	Create(newUser entity.User) (entity.User, error)
	Login(userToLogin dto.LoginUserDto) (entity.User, string, error)
	UpdateLocationById(userToUpdate entity.User) (entity.User, error)
	UpdateDepositById(userId string, addDeposit float64) (entity.User, error)
}
