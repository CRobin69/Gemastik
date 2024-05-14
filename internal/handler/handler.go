package handler

import "github.com/CRobinDev/Gemastik/internal/service"

type Handler struct {
	UserHandler IUserHandler
	LeaderboardHandler ILeaderboardHandler
	ChatHandler IChatHandler
}

func NewHandler(Service *service.Service) *Handler {
	userHandler := NewUserHandler(Service.UserService)
	leaderboardHandler := NewLeaderboardHandler(Service.LeaderboardService)
	chatHandler := NewChatHandler(Service.ChatService)

	return &Handler{
		UserHandler: userHandler,
		LeaderboardHandler: leaderboardHandler,
		ChatHandler: chatHandler,
	}
}
