package chat

import (
	chatv1 "chat/gen/go"
	"chat/internal/domain/models"
	"chat/internal/lib/logger/sl"
	servicechat "chat/internal/services/chat"
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type serverAPI struct {
	chatv1.UnimplementedChatServiceServer
	service servicechat.ChatService
}

func Register(gRPCServer *grpc.Server, service servicechat.ChatService) {
	chatv1.RegisterChatServiceServer(gRPCServer, &serverAPI{
		service: service,
	})
}

func (s *serverAPI) CreateChat(ctx context.Context, req *chatv1.CreateChatRequest) (*chatv1.Chat, error) {
	const op = "grpc.chat.CreateChat"

	if req.GetSenderId() == 0 {
		return nil, status.Error(codes.InvalidArgument, "sender_id is required")
	}
	if req.GetType() == "" {
		return nil, status.Error(codes.InvalidArgument, "type is required")
	}

	ch, err := s.service.Create(ctx, req.GetSenderId(), req.GetInvitedId(), req.GetType(), req.Title)
	if err != nil {
		s.log.Error(op, sl.Err(err))
		return nil, status.Error(codes.Internal, "failed to create chat")
	}

	return s.chatToProto(ctx, ch)
}

func (s *serverAPI) DeleteChat(ctx context.Context, req *chatv1.DeleteChatRequest) (*emptypb.Empty, error) {
	const op = "grpc.chat.DeleteChat"

	if req.GetId() == 0 {
		return nil, status.Error(codes.InvalidArgument, "id is required")
	}

	if err := s.service.Delete(ctx, req.GetId(), req.GetSenderId()); err != nil {
		s.log.Error(op, sl.Err(err))
		return nil, status.Error(codes.Internal, "failed to delete chat")
	}

	return &emptypb.Empty{}, nil
}

func (s *serverAPI) UpdateChat(ctx context.Context, req *chatv1.UpdateChatRequest) (*chatv1.Chat, error) {
	const op = "grpc.chat.UpdateChat"

	if req.GetId() == 0 {
		return nil, status.Error(codes.InvalidArgument, "id is required")
	}

	ch, err := s.service.Update(ctx, req.GetId(), req.Type, req.Title)
	if err != nil {
		s.log.Error(op, sl.Err(err))
		return nil, status.Error(codes.Internal, "failed to update chat")
	}

	return s.chatToProto(ctx, ch)
}

func (s *serverAPI) ListUserChats(ctx context.Context, req *chatv1.ListUserChatsRequest) (*chatv1.ListUserChatsResponse, error) {
	const op = "grpc.chat.ListUserChats"

	if req.GetUserId() == 0 {
		return nil, status.Error(codes.InvalidArgument, "user_id is required")
	}

	chats, err := s.service.ListUserChats(ctx, req.GetUserId())
	if err != nil {
		s.log.Error(op, sl.Err(err))
		return nil, status.Error(codes.Internal, "failed to list chats")
	}

	protoChats := make([]*chatv1.Chat, 0, len(chats))
	for _, ch := range chats {
		protoChat, err := s.chatToProto(ctx, ch)
		if err != nil {
			return nil, status.Error(codes.Internal, "failed to build chat response")
		}
		protoChats = append(protoChats, protoChat)
	}

	return &chatv1.ListUserChatsResponse{Chats: protoChats}, nil
}

func (s *serverAPI) CreateInviteLink(ctx context.Context, req *chatv1.CreateInviteLinkRequest) (*chatv1.InviteLink, error) {
	const op = "grpc.chat.CreateInviteLink"

	if req.GetUserId() == 0 {
		return nil, status.Error(codes.InvalidArgument, "user_id is required")
	}

	link, err := s.service.CreateInviteLink(ctx, req.GetUserId())
	if err != nil {
		s.log.Error(op, sl.Err(err))
		return nil, status.Error(codes.Internal, "failed to create invite link")
	}

	return &chatv1.InviteLink{Url: link}, nil
}

func (s *serverAPI) JoinChatByInvite(ctx context.Context, req *chatv1.JoinChatByInviteRequest) (*chatv1.Chat, error) {
	const op = "grpc.chat.JoinChatByInvite"

	if req.GetId() == 0 {
		return nil, status.Error(codes.InvalidArgument, "id is required")
	}
	if req.GetJoinedId() == 0 {
		return nil, status.Error(codes.InvalidArgument, "joined_id is required")
	}

	ch, err := s.service.JoinChatByInvite(ctx, req.GetId(), req.GetJoinedId())
	if err != nil {
		s.log.Error(op, sl.Err(err))
		return nil, status.Error(codes.Internal, "failed to join chat")
	}

	return s.chatToProto(ctx, ch)
}

func (s *serverAPI) chatToProto(ctx context.Context, ch models.Chat) (*chatv1.Chat, error) {
	users, err := s.service.ChatUsers(ctx, ch.ID)
	if err != nil {
		return nil, err
	}

	adminIDs := make([]int64, 0)
	userIDs := make([]int64, 0)
	for _, u := range users {
		if u.Role == "admin" {
			adminIDs = append(adminIDs, u.UserID)
		}
		userIDs = append(userIDs, u.UserID)
	}

	return &chatv1.Chat{
		Id:        ch.ID,
		Type:      ch.Chat_type,
		Title:     ch.Title,
		CreatedBy: ch.Created_by,
		AdminIds:  adminIDs,
		UserIds:   userIDs,
	}, nil
}
