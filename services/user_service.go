package services

import (
	"time"
	"users/models"
	"users/repositories"
)

type UserService struct {
	repository *repositories.UserRepository
}

func NewUserService(repository *repositories.UserRepository) *UserService {
	return &UserService{repository: repository}
}

func (s *UserService) CreateUser(user *models.User) error {
	return s.repository.CreateUser(user)
}

func (s *UserService) GetUsers() ([]*models.User, error) {
	return s.repository.GetUsers()
}

func (s *UserService) GetUsersByDateAndAge(dateFrom, dateTo time.Time, ageFrom, ageTo int) ([]*models.User, error) {
	return s.repository.GetUsersByDateAndAge(dateFrom, dateTo, ageFrom, ageTo)
}

func (s *UserService) GetUsersCount() (int, error) {
	users, err := s.GetUsers()
	if err != nil {
		return 0, err
	}
	return len(users), nil
}
