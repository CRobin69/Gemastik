package handler

import (
	"net/http"
	"strconv"

	"github.com/CRobinDev/Gemastik/internal/service"
	"github.com/CRobinDev/Gemastik/model"
	"github.com/CRobinDev/Gemastik/pkg/errors"
	"github.com/CRobinDev/Gemastik/pkg/response"
	"github.com/gin-gonic/gin"
)

type ILeaderboardHandler interface {
	CreateLeaderboard(ctx *gin.Context)
	GetLeaderboard(ctx *gin.Context)
	GetLeaderboardByID(ctx *gin.Context)
}

type LeaderboardHandler struct {
	LeaderboardService service.ILeaderboardService
}

func NewLeaderboardHandler(leaderboardService service.ILeaderboardService) ILeaderboardHandler {
	return &LeaderboardHandler{
		LeaderboardService: leaderboardService,
	}
}

func (lh *LeaderboardHandler) CreateLeaderboard(ctx *gin.Context) {
	var req model.CreateLeaderboardRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.Error(ctx, model.ServiceResponse{
			Code:    http.StatusBadRequest,
			Error:   true,
			Message: errors.ErrBadRequest.Error(),
			Data:    nil,
		})
		return
	}

	resp, err := lh.LeaderboardService.CreateLeaderboard(req)
	if err != nil {
		response.Error(ctx, resp)
		return
	}

	response.Success(ctx, resp)
}

func (lh *LeaderboardHandler) GetLeaderboard(ctx *gin.Context) {
	resp, err := lh.LeaderboardService.GetLeaderboard()
	if err != nil {
		response.Error(ctx, model.ServiceResponse{
			Code:    http.StatusInternalServerError,
			Error:   true,
			Message: errors.ErrInternalServer.Error(),
			Data:    nil,
		})
		return
	}

	response.Success(ctx, model.ServiceResponse{
		Code:    http.StatusOK,
		Error:   false,
		Message: "Success",
		Data:    resp,
	})
}

func (lh *LeaderboardHandler) GetLeaderboardByID(ctx *gin.Context) {
	id := ctx.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		response.Error(ctx, model.ServiceResponse{
			Code:    http.StatusBadRequest,
			Error:   true,
			Message: errors.ErrBadRequest.Error(),
			Data:    nil,
		})
		return
	}

	resp, err := lh.LeaderboardService.GetLeaderboardByID(idInt)
	if err != nil {
		response.Error(ctx, model.ServiceResponse{
			Code:    http.StatusInternalServerError,
			Error:   true,
			Message: errors.ErrInternalServer.Error(),
			Data:    err.Error(),
		})
		return
	}

	response.Success(ctx, model.ServiceResponse{
		Code:    http.StatusOK,
		Error:   false,
		Message: "Success",
		Data:    resp,
	})
}
