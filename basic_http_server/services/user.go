package services

import (
	"httpServer/models"
	"httpServer/repositories"
)

type UserService struct {
	UserRepository *repositories.UserRepository
}

func NewUserService(userRepository *repositories.UserRepository) *UserService {
	return &UserService{
		UserRepository: userRepository,
	}
}

func (s *UserService) GetUserByID(id uint) (*models.User, error) {
	user, err := s.UserRepository.FindByID(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UserService) UpdateUser(id uint, profile string) error {
	user, err := s.UserRepository.FindByID(id)
	if err != nil {
		return err
	}
	user.Profile = profile
	return s.UserRepository.Save(user)
}
