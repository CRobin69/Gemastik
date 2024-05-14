package handler

import (
	"net/http"

	"github.com/CRobinDev/Gemastik/entity"
	"github.com/CRobinDev/Gemastik/internal/service"
	"github.com/CRobinDev/Gemastik/model"
	"github.com/CRobinDev/Gemastik/pkg/errors"
	"github.com/CRobinDev/Gemastik/pkg/helper"
	"github.com/CRobinDev/Gemastik/pkg/response"
	"github.com/gin-gonic/gin"
)

type IUserHandler interface {
	RegisterUser(ctx *gin.Context)
	LoginUser(ctx *gin.Context)
	UpdatePassword(ctx *gin.Context)
	UpdateProfile(ctx *gin.Context)
	UploadPhoto(ctx *gin.Context)
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
			Error:   true,
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

	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.Error(ctx, model.ServiceResponse{
			Code:    http.StatusBadRequest,
			Error:   true,
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

func (uh *UserHandler) UpdatePassword(ctx *gin.Context) {
	var req model.UpdatePasswordRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.Error(ctx, model.ServiceResponse{
			Code:    http.StatusBadRequest,
			Error:   true,
			Message: errors.ErrBadRequest.Error(),
			Data:    nil,
		})
		return
	}

	user, err := helper.GetLoginUser(ctx)
	if err != nil {
		response.Error(ctx, model.ServiceResponse{
			Code:    http.StatusUnauthorized,
			Error:   true,
			Message: errors.ErrUnathorized.Error(),
			Data:    err.Error(),
		})
		return
	}
	req.ID = user.Data.(entity.User).ID
	resp, err := uh.UserService.UpdatePassword(req)
	if err != nil {
		response.Error(ctx, resp)
		return
	}

	response.Success(ctx, resp)
}

func (uh *UserHandler) UpdateProfile(ctx *gin.Context) {
	var req model.UserUpdateRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.Error(ctx, model.ServiceResponse{
			Code:    http.StatusBadRequest,
			Error:   true,
			Message: errors.ErrBadRequest.Error(),
			Data:    nil,
		})
		return
	}
	user, err := helper.GetLoginUser(ctx)
	if err != nil {
		response.Error(ctx, model.ServiceResponse{
			Code:    http.StatusUnauthorized,
			Error:   true,
			Message: errors.ErrUnathorized.Error(),
			Data:    err.Error(),
		})
		return
	}
	req.ID = user.Data.(entity.User).ID
	resp, err := uh.UserService.UpdateProfile(req)
	if err != nil {
		response.Error(ctx, resp)
		return
	}

	response.Success(ctx, resp)
}

func (uh *UserHandler) UploadPhoto(ctx *gin.Context) {
	var req model.UploadPhotoRequest

	photo, err := ctx.FormFile("photo")
	if err != nil {
		response.Error(ctx, model.ServiceResponse{
			Code:    http.StatusBadRequest,
			Error:   true,
			Message: errors.ErrBadRequest.Error(),
			Data:    nil,
		})
		return
	}

	user, err := helper.GetLoginUser(ctx)
	if err != nil {
		response.Error(ctx, model.ServiceResponse{
			Code:    http.StatusUnauthorized,
			Error:   true,
			Message: errors.ErrUnathorized.Error(),
			Data:    err.Error(),
		})
		return
	}

	req.ID = user.Data.(entity.User).ID
	req.Photo = photo
	resp, err := uh.UserService.UploadPhoto(req)
	if err != nil {
		response.Error(ctx, resp)
		return
	}

	response.Success(ctx, resp)
}
