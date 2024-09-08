package core

import (
	"context"
	"time"
)

type (
	Chat struct {
		ID        int       `gorm:"column:id"`
		ChatType  string    `gorm:"column:chat_type"`
		IsDeleted bool      `gorm:"column:is_deleted"`
		DeletedAt time.Time `gorm:"column:deleted_at"`
		CreatedAt time.Time `gorm:"column:created_at"`
		UpdatedAt time.Time `gorm:"column:updated_at"`
	}

	Message struct {
		ID        int       `gorm:"column:id"`
		UserID    int       `gorm:"column:user_id"`
		ChatID    int       `gorm:"column:chat_id"`
		Content   string    `gorm:"column:content"`
		IsDeleted bool      `gorm:"column:is_deleted"`
		DeletedAt time.Time `gorm:"column:deleted_at"`
		CreatedAt time.Time `gorm:"column:created_at"`
		UpdatedAt time.Time `gorm:"column:updated_at"`
	}

	ChatMember struct {
		ID        int       `gorm:"column:id"`
		UserID    int       `gorm:"column:user_id"`
		ChatID    int       `gorm:"column:chat_id"`
		IsDeleted bool      `gorm:"column:is_deleted"`
		DeletedAt time.Time `gorm:"column:deleted_at"`
		CreatedAt time.Time `gorm:"column:created_at"`
	}

	ChatStore interface {
		GetAllChats(ctx context.Context, sortType string) (chats []Chat, err error)
		GetChatByID(ctx context.Context, id int) (chat Chat, err error)
		CreateChat(ctx context.Context, data Chat) (chat Chat, err error)
		DeleteChat(ctx context.Context, id int) (err error)
	}

	ChatService interface {
		GetAllChats(ctx context.Context, sortType string) (chats []Chat, total int, err error)
		GetChatByID(ctx context.Context, id int) (chat Chat, err error)
		CreateChat(ctx context.Context, data Chat) (chat Chat, err error)
		DeleteChat(ctx context.Context, id int) (err error)
	}

	MessageStore interface {
		GetAllMessages(ctx context.Context, chatID int, sortType string, searchText string) (messages []Message, err error)
		CreateMessage(ctx context.Context, data Message) (message Message, err error)
		UpdateMessage(ctx context.Context, chatID int, messageID int, data string) (message Message, err error)
		DeleteMessage(ctx context.Context, chatID int, messageID int) (err error)
	}

	MessageService interface {
		GetAllMessages(ctx context.Context, chatID int, sortType string, searchText string) (messages []Message, total int, err error)
		CreateMessage(ctx context.Context, data Message) (message Message, err error)
		UpdateMessage(ctx context.Context, chatID int, messageID int, data string) (message Message, err error)
		DeleteMessage(ctx context.Context, chatID int, messageID int) (err error)
	}

	ChatMemberStore interface {
		GetAllMembers(ctx context.Context, chatID int) (members []ChatMember, err error)
		AddMemberToChat(ctx context.Context, data ChatMember) (member ChatMember, err error)
		UpdateMemberInfo(ctx context.Context, chatID int, userID int) (member ChatMember, err error)
		DeleteMemberFromChat(ctx context.Context, chatID int, userID int) (err error)
	}

	ChatMemberService interface {
		GetAllMembers(ctx context.Context, chatID int) (members []ChatMember, total int, err error)
		AddMemberToChat(ctx context.Context, data ChatMember) (member ChatMember, err error)
		UpdateMemberInfo(ctx context.Context, chatID int, userID int) (member ChatMember, err error)
		DeleteMemberFromChat(ctx context.Context, chatID int, userID int) (err error)
	}
)