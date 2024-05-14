package entity

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID              uuid.UUID `json:"id" gorm:"type:varchar(36);primary_key;"`
	PhoneNumber     string    `json:"phone_number" gorm:"type:varchar(13);not null;"`
	Name            string    `json:"name" gorm:"type:varchar(255);not null;"`
	Email           string    `json:"email" gorm:"type:varchar(255);not null;unique"`
	Password        string    `json:"password" gorm:"type:varchar(255);not null;"`
	ConfirmPassword string    `json:"confirm_password" gorm:"-"`
	PhotoLink       string    `json:"photo_link" gorm:"type:text"`
	CreatedAt       time.Time `json:"createdAt" gorm:"autoCreateTime"`
	UpdatedAt       time.Time `json:"updatedAt" gorm:"autoUpdateTime"`
}
