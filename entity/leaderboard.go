package entity

import "time"

type Leaderboard struct {
	ID              int       `json:"id" gorm:"primaryKey;autoIncrement"`
	Name            string    `json:"name" gorm:"type:varchar(255);not null;"`
	PhotoLink       string    `json:"photo_link" gorm:"type:text"`
	TotalCorrupt    string    `json:"total_corruption" gorm:"type:varchar(36);not null;"`
	DetailCorruptor string    `json:"detail_corruptor" gorm:"type:text;not null;"`
	CreatedAt       time.Time `json:"createdAt"`
	UpdatedAt       time.Time `json:"updatedAt"`
}
