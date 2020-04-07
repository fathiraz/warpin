package usecase

import (
	"context"
	"warpin/configs"
	"warpin/internal/chats/domain"
	"warpin/internal/chats/repository"
	"warpin/tools"
)

type chatUsecase struct {
	repository *repository.ChatRepository
}

// NewChatsUsecase to implement chat usecase interface
func NewChatsUsecase(repository *repository.ChatRepository) ChatUsecase {
	return &chatUsecase{
		repository: repository,
	}
}

func (c *chatUsecase) Insert(context context.Context, data *domain.Chats) (*domain.Chats, error) {
	ctx := "ChatUsecase-Insert"

	var (
		response *domain.Chats
		err      error
	)

	// insert data
	createRepo := <-c.repository.Chats.Create(context, data)
	if createRepo.Error != nil {
		err = createRepo.Error
		tools.Log(tools.ErrorLevel, err.Error(), ctx, "create_chats_repository")
		return response, err
	}

	// transform interface to struct
	response = createRepo.Data.(*domain.Chats)

	// return response
	return response, nil
}

func (c *chatUsecase) FindAll(context context.Context, filter *domain.Filter) ([]*domain.Chats, error) {
	ctx := "ChatUsecase-FindAll"

	var (
		response []*domain.Chats
		err      error
	)

	// get data
	findRepo := <-c.repository.Chats.GetAll(context, filter)
	if findRepo.Error != nil {
		err = findRepo.Error
		tools.Log(tools.ErrorLevel, err.Error(), ctx, "getall_chats_repository")
		return response, err
	}

	// transform interface to struct
	response = findRepo.Data.([]*domain.Chats)

	// return response
	return response, nil
}

func (c *chatUsecase) Websocket(context context.Context, name string) {
	ctx := "ChatUsecase-Websocket"

	var (
		// for broadcast message
		broadcastMessage = &domain.ChatSocketResponse{}

		// payload request
		chatRequest = &domain.ChatSocketRequest{}
	)

	// set broadcast message for new name
	broadcastMessage = &domain.ChatSocketResponse{
		Name: name,
		Type: domain.ChatComeType,
	}

	// broadcast message from a new name
	configs.Broadcast(name, broadcastMessage)

	// loop to check user typing message
	for {
		// get ws connections
		ws := configs.GetConnectionsByName(name)

		// check ws not null
		if ws != nil {

			// read current connection to our payload
			err := ws.Conn.ReadJSON(&chatRequest)
			if err != nil {
				tools.Log(tools.ErrorLevel, err.Error(), ctx, "websocket_chat")
				break
			}

			// set broadcast message for new chat
			broadcastMessage = &domain.ChatSocketResponse{
				Name:    name,
				Type:    domain.ChatMessageType,
				Message: chatRequest.Message,
			}

			// broadcast message from a new chat
			configs.Broadcast(name, broadcastMessage)
		}
	}
}
