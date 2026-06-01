package chat

import (
	"context"

	"github.com/PoinkCorporation/Chat-service/internal/domain/models"
	"github.com/PoinkCorporation/Chat-service/internal/storage/postgresql"
)

type PermissionsChecker interface {
	HasPermission(ctx context.Context, userID int64, permission string) (bool, error)
}

type chat struct {
	storage     *postgresql.PGRepo
	permChecker PermissionsChecker
}

func New(storage *postgresql.PGRepo) *chat {
	return &chat{storage: storage}
}

func (s *chat) Create(ctx context.Context, senderID int64, invitedIDs []int64, chatType string, title *string) (models.Chat, error) {
	ch := models.Chat{
		ChatType:  chatType,
		Title:     title,
		CreatedBy: senderID,
	}

	createdChat, err := s.storage.NewChat(ctx, ch)
	if err != nil {
		return models.Chat{}, err
	}

	if err := s.storage.AddChatUser(ctx, senderID, createdChat.ID, "admin"); err != nil {
		return models.Chat{}, err
	}

	for _, invitedID := range invitedIDs {
		if err := s.storage.AddChatUser(ctx, invitedID, createdChat.ID, "member"); err != nil {
			return models.Chat{}, err
		}
	}

	return createdChat, nil
}

func (s *chat) Delete(ctx context.Context, id, senderID int64) error {
	return s.storage.DeleteChat(ctx, id)
}

func (s *chat) Update(ctx context.Context, id int64, chatType, title *string) (models.Chat, error) {
	currentChat, err := s.storage.ChatByID(ctx, id)
	if err != nil {
		return models.Chat{}, err
	}

	if chatType != nil {
		currentChat.ChatType = *chatType
	}
	if title != nil {
		currentChat.Title = title
	}

	if err := s.storage.UpdateChat(ctx, currentChat); err != nil {
		return models.Chat{}, err
	}

	return currentChat, nil
}

func (s *chat) ListUserChats(ctx context.Context, userID int64) ([]models.Chat, error) {
	return s.storage.ListUserChats(ctx, userID)
}

func (s *chat) CreateInviteLink(ctx context.Context, userID int64) (string, error) {
	return "", nil
}

func (s *chat) JoinChatByInvite(ctx context.Context, chatID, joinedID int64) (models.Chat, error) {
	return models.Chat{}, nil
}

func (s *chat) ChatUsers(ctx context.Context, chatID int64) ([]models.ChatUser, error) {
	return s.storage.ChatUsers(ctx, chatID)
}
