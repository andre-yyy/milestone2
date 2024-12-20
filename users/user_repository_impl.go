package users

import (
	"fmt"

	entity "milestone2/entities"
	repository "milestone2/repositories"
	dto "milestone2/users/dtos"
	utility "milestone2/utilities"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type userRepositoryImpl struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) repository.UserRepository {
	return &userRepositoryImpl{DB: db}
}

func (userRepository *userRepositoryImpl) Find() ([]entity.User, error) {
	var user []entity.User
	err := userRepository.DB.Find(&user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (userRepository *userRepositoryImpl) FindOneById(userId string) (entity.User, error) {
	var user entity.User
	err := userRepository.DB.Where(&entity.User{ID: userId}).Take(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound || len(user.ID) <= 0 {
			return entity.User{}, fmt.Errorf("user with id '%s' not found", userId)
		}
		return entity.User{}, err
	}
	return user, nil
}

func (userRepository *userRepositoryImpl) FindOneByEmail(userEmail string) (entity.User, error) {
	var user entity.User
	err := userRepository.DB.Where(&entity.User{Email: userEmail}).Take(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound || len(user.ID) <= 0 {
			return entity.User{}, fmt.Errorf("user with email '%s' not found", userEmail)
		}
		return entity.User{}, err
	}
	return user, nil
}

func (userRepository *userRepositoryImpl) Create(newUser entity.User) (entity.User, error) {
	user, _ := userRepository.FindOneByEmail(newUser.Email)
	if len(user.ID) <= 0 {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
		if err != nil {
			return entity.User{}, err
		}
		newUser.Password = string(hashedPassword)
		err = userRepository.DB.Create(&newUser).Error
		if err != nil {
			return entity.User{}, fmt.Errorf("failed to create user : '%s'", err)
		}
		return newUser, nil
	}
	return entity.User{}, fmt.Errorf("email already registered")
}

func (userRepository *userRepositoryImpl) Login(userToLogin dto.LoginUserDto) (entity.User, string, error) {
	user, err := userRepository.FindOneByEmail(userToLogin.Email)
	if len(user.ID) <= 0 {
		return entity.User{}, "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userToLogin.Password))
	if err != nil {
		return entity.User{}, "", fmt.Errorf("incorrect password")
	}
	tokenString, err := utility.GenerateToken(user.ID)
	if err != nil {
		return entity.User{}, "", err
	}
	return user, tokenString, nil
}

func (userRepository *userRepositoryImpl) UpdateLocationById(userToUpdate entity.User) (entity.User, error) {
	user, err := userRepository.FindOneById(userToUpdate.ID)
	if err != nil {
		return entity.User{}, err
	}
	user.Address = userToUpdate.Address
	user.CityID = userToUpdate.CityID
	err = userRepository.DB.Save(&user).Error
	if err != nil {
		return entity.User{}, err
	}
	return user, nil
}

func (userRepository *userRepositoryImpl) UpdateDepositById(userId string, addDeposit float64) (entity.User, error) {
	user, err := userRepository.FindOneById(userId)
	if err != nil {
		return entity.User{}, err
	}
	user.Deposit += addDeposit
	err = userRepository.DB.Save(&user).Error
	if err != nil {
		return entity.User{}, err
	}
	return user, nil
}
