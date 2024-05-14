package model

import "github.com/google/uuid"

type ChatRequest struct {
	UserID uuid.UUID `json:"-"`
	Chat   string `json:"chat" binding:"required"`
}
