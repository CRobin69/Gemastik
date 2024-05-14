package repository

import (
	"github.com/CRobinDev/Gemastik/entity"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type IChatRepository interface {
	InsertChat(chat entity.Chat) error
	GetHistory(userID uuid.UUID) ([]string, error)
}

type ChatRepository struct {
	db *gorm.DB
}

func NewChatRepository(db *gorm.DB) IChatRepository {
	return &ChatRepository{
		db: db,
	}
}

func (r *ChatRepository) InsertChat(chat entity.Chat) error {
	if err := r.db.Create(&chat).Error; err != nil {
		return err
	}

	return nil
}

func (r *ChatRepository) GetHistory(userID uuid.UUID) ([]string, error) {
    var inputs []string

    if err := r.db.Model(&entity.Chat{}).
        Where("user_id = ?", userID).
        Pluck("input", &inputs).Error; err != nil {
        return nil, err
    }

    return inputs, nil
}
