package repository

import (
	"github.com/CRobinDev/Gemastik/entity"
	"github.com/CRobinDev/Gemastik/model"
	"gorm.io/gorm"
)

type ILeaderboardRepository interface {
	CreateLeaderboard(leaderboard entity.Leaderboard) error
	GetLeaderboard() ([]model.CreateLeaderboardResponse, error)
	GetLeaderboardByID(id int) (model.CreateLeaderboardResponse, error)
}

type LeaderboardRepository struct {
	db *gorm.DB
}

func NewLeaderboardRepository(db *gorm.DB) ILeaderboardRepository {
	return &LeaderboardRepository{db: db}
}

func (r *LeaderboardRepository) CreateLeaderboard(leaderboard entity.Leaderboard) error {
	if err := r.db.Create(&leaderboard).Error; err != nil {
		return err
	}
	return nil
}

func (r *LeaderboardRepository) GetLeaderboard() ([]model.CreateLeaderboardResponse, error) {
	var leaderboards []entity.Leaderboard

	if err := r.db.Find(&leaderboards).Error; err != nil {
		return nil, err
	}

	var response []model.CreateLeaderboardResponse
	for _, leaderboard := range leaderboards {
		response = append(response, model.CreateLeaderboardResponse{
			ID:          leaderboard.ID,
			Name: 	  leaderboard.Name,
			PhotoLink:    leaderboard.PhotoLink,
			TotalCorrupt: leaderboard.TotalCorrupt,
			DetailCorruptor: leaderboard.DetailCorruptor,
			CreatedAt:    leaderboard.CreatedAt,

		})
	}

	return response, nil
}

func (r *LeaderboardRepository) GetLeaderboardByID(id int) (model.CreateLeaderboardResponse, error) {
	var leaderboard entity.Leaderboard

	if err := r.db.Where("id = ?", id).First(&leaderboard).Error; err != nil {
		return model.CreateLeaderboardResponse{}, err
	}

	return model.CreateLeaderboardResponse{
		ID:          leaderboard.ID,
		Name: 	  leaderboard.Name,
		PhotoLink:    leaderboard.PhotoLink,
		TotalCorrupt: leaderboard.TotalCorrupt,
		DetailCorruptor: leaderboard.DetailCorruptor,
		CreatedAt:    leaderboard.CreatedAt,
	}, nil
}
