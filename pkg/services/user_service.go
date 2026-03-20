package services

import (
	"example.com/project/pkg/models"
	"example.com/project/pkg/repository"
)

// UserService provides user related functionality
// @Description user service
// @receiver us
// @return UserService
func NewUserService() *UserService {
	return &UserService{UserRepo: repository.NewUserRepository()}
}

// GetUserSvc provides user svc instance
func GetUserSvc() *UserService {
	return NewUserService()
}

type UserService struct {
	UserRepo repository.UserRepository
}

func (us *UserService) GetUsers() ([]models.User, error) {
	return us.UserRepo.FindAllUsers()
}