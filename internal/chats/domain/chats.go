package domain

import (
	"time"
	"warpin/tools"
)

const (
	TableName = "chats"

	// const for broadcast message
	ChatComeType    = "memasuki chatroom"
	ChatMessageType = "chat"
)

type Chats struct {
	ID        int       `gorm:"column:id;AUTO_INCREMENT" json:"id"`
	Name      string    `gorm:"column:name" json:"name" validate:"required"`
	Message   string    `gorm:"column:message" json:"message" validate:"required"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
}

func (c *Chats) TableName() string {
	return TableName
}

type Filter struct {
	tools.Filter
}

type ChatSocketRequest struct {
	Message string
}

type ChatSocketResponse struct {
	Name    string
	Type    string
	Message string
}
