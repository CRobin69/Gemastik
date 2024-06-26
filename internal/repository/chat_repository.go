package repository

import (
	"log"
	"time"

	"github.com/CRobinDev/Gemastik/entity"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type IChatRepository interface {
	InsertChat(chat entity.Chat) error
	GetHistory(userID uuid.UUID) ([]string, error)
	DeleteOldChats() error
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

func (r *ChatRepository) DeleteOldChats() error {
	loc, _ := time.LoadLocation("Asia/Jakarta")
	createdAt, _ := time.Parse("2006-01-02 15:04:05", time.Now().In(loc).Format("2006-01-02 15:04:05"))

	deleteTime := createdAt.Add(-24 * time.Hour) 

	log.Printf("Current time: %v\n", createdAt)
	log.Printf("Delete records before time : %v\n", deleteTime)

	var count int64
	r.db.Model(&entity.Chat{}).Where("created_at < ?", deleteTime).Count(&count)
	log.Printf("Number of records to delete: %d\n", count)

	if err := r.db.Where("created_at < ?", deleteTime).Delete(&entity.Chat{}).Error; err != nil {
		log.Printf("Error deleting old chats: %v\n", err)
		return err
	}

	r.db.Model(&entity.Chat{}).Count(&count)
	log.Printf("Number of records remaining: %d\n", count)

	return nil
}
