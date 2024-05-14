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

type IChatHandler interface {
	GenerateResponse(ctx *gin.Context)
}

type ChatHandler struct {
	ChatService service.IChatService
}

func NewChatHandler(chatService service.IChatService) IChatHandler {
	return &ChatHandler{
		ChatService: chatService,
	}
}

func (ch *ChatHandler) GenerateResponse(ctx *gin.Context) {
	var req model.ChatRequest
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
	req.UserID = user.Data.(entity.User).ID

	resp, err := ch.ChatService.GenerateResponse(req)
	if err != nil {
		response.Error(ctx, resp)
		return
	}

	response.Success(ctx, resp)
}
