package service

import (
	"github.com/CRobinDev/Gemastik/internal/repository"
)

type Service struct {
	UserService IUserService
}

func NewService(Repository *repository.Repository) *Service {
	userService := NewUserService(Repository.UserRepository)

	return &Service{
		UserService: userService,
	}
}
