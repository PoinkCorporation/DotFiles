package chat

import (
	"context"

	chatv1 "chat/gen/go"
	"chat/internal/domain/models"

	"google.golang.org/grpc"
)

type serverAPI struct {
	chatv1.UnimplementedChatServiceServer
	chatService ChatService
}

type ChatService interface {
	CreateChat(
		ctx context.Context,
		senderID int64,
		invitedIDs []int64,
		chatType string,
		title *string,
	) (models.Chat, error)
	DeleteChat(
		ctx context.Context,
		id,
		senderID int64,
	) error
	UpdateChat(
		ctx context.Context,
		id int64,
		chatType,
		title *string,
	) (models.Chat, error)
	ListUserChats(
		ctx context.Context,
		userID int64,
	) ([]models.Chat, error)
	CreateInviteLink(
		ctx context.Context,
		userID int64,
	) (string, error)
	JoinChatByInvite(
		ctx context.Context,
		chatID,
		joinedID int64,
	) (models.Chat, error)
}

func Register(grpcServer *grpc.Server, ChatService ChatService) {
	chatv1.RegisterChatServiceServer(grpcServer, serverAPI{chatService: ChatService})
}
