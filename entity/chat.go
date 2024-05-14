package entity

import (
	"time"

	"github.com/google/uuid"
)

type Chat struct {
	ID        int       `json:"id" gorm:"primaryKey;autoIncrement"`
	UserID    uuid.UUID `json:"userid" gorm:"type:varchar(36);foreignkey:ID;references:users;onUpdate:CASCADE;onDelete:CASCADE"`
	Input     string    `json:"input" gorm:"type:text"`
	Output    string    `json:"output" gorm:"type:text"`
	CreatedAt time.Time `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"autoUpdateTime"`
}
