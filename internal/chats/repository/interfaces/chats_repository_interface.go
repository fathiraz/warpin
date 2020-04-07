package interfaces

import (
	"context"
	"warpin/internal/chats/domain"
	"warpin/tools"
)

type ChatsRepository interface {
	Create(context context.Context, newData *domain.Chats) <-chan tools.RepositoryResult
	GetAll(context context.Context, filter *domain.Filter) <-chan tools.RepositoryResult
}
