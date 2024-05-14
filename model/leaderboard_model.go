package model

import "time"

type CreateLeaderboardRequest struct {
	Name            string `json:"name" binding:"required"`
	PhotoLink       string `json:"photo_link"`
	TotalCorrupt    string `json:"total_corruption" binding:"required"`
	DetailCorruptor string `json:"detail_corruptor" binding:"required"`
}

type CreateLeaderboardResponse struct {
	ID              int       `json:"id"`
	Name            string    `json:"name"`
	PhotoLink       string    `json:"photo_link"`
	TotalCorrupt    string    `json:"total_corruption"`
	DetailCorruptor string    `json:"detail_corruptor"`
	CreatedAt        time.Time `json:"created_at"`
}
