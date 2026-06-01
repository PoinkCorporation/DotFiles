package app

import (
	grpcapp "chat/internal/app/grpc"
	"chat/internal/services/chat"
	"chat/internal/storage/postgresql"
	"log/slog"
)

type App struct {
	GRPCServer *grpcapp.App
}

func New(
	log *slog.Logger,
	grpcPort int,
	storagePath string,
) *App {
	storage, err := postgresql.New(storagePath)
	if err != nil {
		panic(err)
	}

	chatService := chat.New(storage)

	grpcApp := grpcapp.New(log, chatService, grpcPort)

	return &App{
		GRPCServer: grpcApp,
	}
}
