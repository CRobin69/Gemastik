package handler

import (
	"net/http"

	"github.com/CRobinDev/Gemastik/internal/service"
	"github.com/CRobinDev/Gemastik/model"
	"github.com/CRobinDev/Gemastik/pkg/errors"
	"github.com/CRobinDev/Gemastik/pkg/response"
	"github.com/gin-gonic/gin"
)

type IUserHandler interface {
	RegisterUser(ctx *gin.Context)
	LoginUser(ctx *gin.Context)
}

type UserHandler struct {
	UserService service.IUserService
}

func NewUserHandler(userService service.IUserService) IUserHandler {
	return &UserHandler{
		UserService: userService,
	}
}

func (uh *UserHandler) RegisterUser(ctx *gin.Context) {
	var req model.UserRegisterRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.Error(ctx, model.ServiceResponse{
			Code:    http.StatusBadRequest,
			Error : true,
			Message: errors.ErrBadRequest.Error(),
			Data:    nil,
		})
		return
	}

	resp, err := uh.UserService.Register(req)
	if err != nil {
		response.Error(ctx, resp)
		return
	}

	response.Success(ctx, resp)
}

func (uh *UserHandler) LoginUser(ctx *gin.Context) {
	var req model.UserLoginRequest

	if err := ctx.ShouldBindJSON(&req);err != nil {
		response.Error(ctx, model.ServiceResponse{
			Code:    http.StatusBadRequest,
			Error : true,
			Message: errors.ErrBadRequest.Error(),
			Data:    nil,
		})
		return
	}

	resp, err := uh.UserService.LoginUser(req)
	if err != nil {
		response.Error(ctx, resp)
		return
	}

	response.Success(ctx, resp)
}