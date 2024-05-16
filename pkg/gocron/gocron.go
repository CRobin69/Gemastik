package gocron

import (
	"log"
	"time"

	"github.com/CRobinDev/Gemastik/internal/repository"
	"github.com/go-co-op/gocron"
)

func StartChatCleanupScheduler(chatRepo repository.IChatRepository) {
	s := gocron.NewScheduler(time.UTC)
	s.Every(2).Hour().Do(func() {
		err := chatRepo.DeleteOldChats()
		if err != nil {
			log.Printf("Error cleaning up old chats: %v", err)
		} else {
			log.Println("Successfully cleaned up old chats")
		}
	})
	s.StartAsync()
}
