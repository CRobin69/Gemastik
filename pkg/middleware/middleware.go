package middleware

import (
	"github.com/CRobinDev/Gemastik/internal/service"
	"github.com/gin-gonic/gin"
)

type Interface interface {
	AuthenticateUser(ctx *gin.Context)
}

type middleware struct {
	service *service.Service
}

func Init(service *service.Service) Interface {
	return &middleware{
		service: service,
	}
}
