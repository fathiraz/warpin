package sqlite

import (
	"context"
	"fmt"
	"github.com/jinzhu/gorm"
	"time"
	"warpin/internal/chats/domain"
	"warpin/internal/chats/repository/interfaces"
	"warpin/tools"
)

type chatsRepositorySqlite struct {
	db *gorm.DB
}

// NewChatsRepositorySqlite function to implement repository interfaces
func NewChatsRepositorySqlite(db *gorm.DB) interfaces.ChatsRepository {
	// auto migrate model when empty
	db.AutoMigrate(&domain.Chats{})

	return &chatsRepositorySqlite{db: db}
}

// Create function to create chats
func (c *chatsRepositorySqlite) Create(context context.Context, newData *domain.Chats) <-chan tools.RepositoryResult {
	result := make(chan tools.RepositoryResult)

	go func() {
		// set our db
		db := c.db

		newData.CreatedAt = time.Now()

		// insert data
		if err := db.Create(&newData).Error; err != nil {
			result <- tools.RepositoryResult{Error: err}
			return
		}

		// set result
		result <- tools.RepositoryResult{Data: newData}
	}()

	return result
}

// GetAll function to get all chats
func (c *chatsRepositorySqlite) GetAll(context context.Context, filter *domain.Filter) <-chan tools.RepositoryResult {
	result := make(chan tools.RepositoryResult)

	go func() {
		// set our db
		db := c.db

		// set order by
		orderBy := fmt.Sprintf("%s %s", filter.OrderBy, filter.Sort)

		// get data
		var datas []*domain.Chats
		if err := db.Limit(filter.Limit).Offset(filter.Offset).Order(orderBy).Find(&datas).Error; err != nil {
			result <- tools.RepositoryResult{Error: err}
			return
		}

		// set result
		result <- tools.RepositoryResult{Data: datas}
	}()

	return result
}
