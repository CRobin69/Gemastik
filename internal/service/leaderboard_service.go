package service

import (
	"net/http"
	"time"

	"github.com/CRobinDev/Gemastik/entity"
	"github.com/CRobinDev/Gemastik/internal/repository"
	"github.com/CRobinDev/Gemastik/model"
	"github.com/CRobinDev/Gemastik/pkg/errors"
)

type ILeaderboardService interface {
	CreateLeaderboard(request model.CreateLeaderboardRequest) (model.ServiceResponse, error)
	GetLeaderboard() ([]model.ServiceResponse, error)
	GetLeaderboardByID(id int) (model.ServiceResponse, error)
}

type LeaderboardService struct {
	LeaderboardRepository repository.ILeaderboardRepository
}

func NewLeaderboardService(leaderboardRepository repository.ILeaderboardRepository) ILeaderboardService {
	return &LeaderboardService{
		LeaderboardRepository: leaderboardRepository,
	}
}

func (ls *LeaderboardService) CreateLeaderboard(request model.CreateLeaderboardRequest) (model.ServiceResponse, error) {
	loc, _ := time.LoadLocation("Asia/Jakarta")
	createdAt, _ := time.Parse("2006-01-02 15:04:05", time.Now().In(loc).Format("2006-01-02 15:04:05"))
	leaderboard := entity.Leaderboard{
		Name:            request.Name,
		PhotoLink:       request.PhotoLink,
		TotalCorrupt:    request.TotalCorrupt,
		DetailCorruptor: request.DetailCorruptor,
		CreatedAt:       createdAt,
	}

	if err := ls.LeaderboardRepository.CreateLeaderboard(leaderboard); err != nil {
		return model.ServiceResponse{
			Code:    http.StatusInternalServerError,
			Error:   true,
			Message: errors.ErrFailedCreateBoard.Error(),
			Data:    nil,
		}, err
	}

	return model.ServiceResponse{
		Code:    http.StatusOK,
		Error:   false,
		Message: "Leaderboard created successfully",
		Data:    nil,
	}, nil
}

func (ls *LeaderboardService) GetLeaderboard() ([]model.ServiceResponse, error) {
	leaderboards, err := ls.LeaderboardRepository.GetLeaderboard()
	if err != nil {
		return []model.ServiceResponse{
			{
				Code:    http.StatusInternalServerError,
				Error:   true,
				Message: errors.ErrInternalServer.Error(),
				Data:    nil,
			},
		}, err
	}

	var response []model.ServiceResponse
	for _, leaderboard := range leaderboards {
		response = append(response, model.ServiceResponse{
			Code:    http.StatusOK,
			Error:   false,
			Message: "Success",
			Data:    leaderboard,
		})
	}

	return response, nil
}

func (ls *LeaderboardService) GetLeaderboardByID(id int) (model.ServiceResponse, error) {
	leaderboard, err := ls.LeaderboardRepository.GetLeaderboardByID(id)
	if err != nil {
		return model.ServiceResponse{
			Code:    http.StatusInternalServerError,
			Error:   true,
			Message: errors.ErrInternalServer.Error(),
			Data:    nil,
		}, err
	}

	return model.ServiceResponse{
		Code:    http.StatusOK,
		Error:   false,
		Message: "Success",
		Data:    leaderboard,
	}, nil
}
