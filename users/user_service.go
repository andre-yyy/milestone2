package users

import (
	entity "milestone2/entities"
	repository "milestone2/repositories"
	dto "milestone2/users/dtos"
)

type UserService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) *UserService {
	return &UserService{userRepository: userRepository}
}

func (s *UserService) GetUserById(userId string) (entity.User, error) {
	user, err := s.userRepository.FindOneById(userId)
	if err != nil {
		return entity.User{}, err
	}
	return user, nil
}

func (s *UserService) GetUserByEmail(userEmail string) (entity.User, error) {
	user, err := s.userRepository.FindOneByEmail(userEmail)
	if err != nil {
		return entity.User{}, err
	}
	return user, nil
}

func (s *UserService) LoginUser(userToLogin dto.LoginUserDto) (entity.User, string, error) {
	loggedInUser, tokenString, err := s.userRepository.Login(userToLogin)
	if err != nil {
		return entity.User{}, "", err
	}
	return loggedInUser, tokenString, nil
}

func (s *UserService) RegisterUser(newUser entity.User) (entity.User, error) {
	registeredUser, err := s.userRepository.Create(newUser)
	if err != nil {
		return entity.User{}, err
	}
	return registeredUser, nil
}

func (s *UserService) UpdateUserLocationById(userToUpdate entity.User) (entity.User, error) {
	updatedUser, err := s.userRepository.UpdateLocationById(userToUpdate)
	if err != nil {
		return entity.User{}, err
	}
	return updatedUser, nil
}

func (s *UserService) UpdateUserDepositById(userId string, addDeposit float64) (entity.User, error) {
	updatedUser, err := s.userRepository.UpdateDepositById(userId, addDeposit)
	if err != nil {
		return entity.User{}, err
	}
	return updatedUser, nil
}
