package middleware

import (
	"github.com/CRobinDev/Gemastik/internal/service"
	"github.com/gin-gonic/gin"
)

type IMiddleware interface {
	AuthenticateUser(ctx *gin.Context)
}

type Middleware struct {
	service *service.Service
}

func Init(service *service.Service) IMiddleware {
	return &Middleware{
		service: service,
	}
}
