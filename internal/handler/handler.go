package handler

import "github.com/CRobinDev/Gemastik/internal/service"

type Handler struct {
	UserHandler IUserHandler
}

func NewHandler(Service *service.Service) *Handler {
	userHandler := NewUserHandler(Service.UserService)

	return &Handler{
		UserHandler: userHandler,
	}
}
