package services

import (
  db "user-management/internal/database"
  "user-management/internal/models"
)

type UserService struct {
  userRepository db.UserRepository
}

func NewUserService(userRepository db.UserRepository) *UserService {
  return &UserService{userRepository: userRepository}
}

func (us *UserService) RegisterUser(user *models.User) error {
  return us.userRepository.RegisterUser(user)
}

func (us *UserService) GetUser(id int) (*models.User, error) {
  return us.userRepository.GetUserById(id)
}

func (us *UserService) DeleteUser(id int) error {
  return us.userRepository.DeleteUser(id)
}
