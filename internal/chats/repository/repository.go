package repository

import (
	"github.com/jinzhu/gorm"
	"warpin/internal/chats/repository/interfaces"
	"warpin/internal/chats/repository/sqlite"
)

// ChatRepository struct to hold our repository
type ChatRepository struct {
	Chats interfaces.ChatsRepository
}

// NewChatsRepository for initiate chats repository
func NewChatsRepository(db *gorm.DB) *ChatRepository {
	return &ChatRepository{
		Chats: sqlite.NewChatsRepositorySqlite(db),
	}
}
