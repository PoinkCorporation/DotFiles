package chat

import (
	"context"

	chatv1 "chat/gen/go"
	"chat/internal/domain/models"
)

type serverAPI struct {
	chatv1.UnimplementedChatServiceServer
}

type ChatService interface {
	CreateChat(
		ctx context.Context,
		senderID int64,
		invitedIDs []int64,
		chatType string,
		title *string,
	) (models.Chat, error)
	DeleteChat(ctx context.Context, id, senderID int64) error
	UpdateChat(ctx context.Context, id int64, chatType, title *string) (models.Chat, error)
	ListUserChats(ctx context.Context, userID int64) ([]models.Chat, error)
	CreateInviteLink(ctx context.Context, userID int64) (string, error)
	JoinChatByInvite(ctx context.Context, chatID, joinedID int64) (models.Chat, error)
}
