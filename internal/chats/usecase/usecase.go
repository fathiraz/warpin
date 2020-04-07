package usecase

import (
	"context"
	"warpin/internal/chats/domain"
)

type ChatUsecase interface {
	Insert(context context.Context, data *domain.Chats) (*domain.Chats, error)
	FindAll(context context.Context, filter *domain.Filter) ([]*domain.Chats, error)
	Websocket(context context.Context, name string)
}
