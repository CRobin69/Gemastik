package repository

import "gorm.io/gorm"

type Repository struct {
	UserRepository IUserRepository
	LeaderboardRepository ILeaderboardRepository
	ChatRepository IChatRepository
}

func NewRepository(db *gorm.DB) *Repository {
	userRepository := NewUserRepository(db)
	leaderboardRepository := NewLeaderboardRepository(db)
	chatRepository := NewChatRepository(db)

	return &Repository{
		UserRepository: userRepository,
		LeaderboardRepository: leaderboardRepository,
		ChatRepository: chatRepository,
	}
}
