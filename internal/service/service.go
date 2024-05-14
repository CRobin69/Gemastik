package service

import (
	"github.com/CRobinDev/Gemastik/internal/repository"
)

type Service struct {
	UserService        IUserService
	LeaderboardService ILeaderboardService
	ChatService        IChatService
}

func NewService(Repository *repository.Repository) *Service {
	userService := NewUserService(Repository.UserRepository)
	leaderboardService := NewLeaderboardService(Repository.LeaderboardRepository)
	chatService := NewChatService(Repository.UserRepository, Repository.ChatRepository)

	return &Service{
		UserService:        userService,
		LeaderboardService: leaderboardService,
		ChatService:        chatService,
	}
}
