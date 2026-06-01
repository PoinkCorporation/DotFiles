package app

import (
	"log/slog"

	grpcapp "github.com/PoinkCorporation/Chat-service/internal/app/grpc"
	clients "github.com/PoinkCorporation/Chat-service/internal/clients/permissions"
	"github.com/PoinkCorporation/Chat-service/internal/services/chat"
	"github.com/PoinkCorporation/Chat-service/internal/storage/postgresql"
)

type App struct {
	GRPCServer *grpcapp.App
}

func New(
	log *slog.Logger,
	grpcPort int,
	storagePath string,
	permissionsPath string,
) *App {
	storage, err := postgresql.New(storagePath)
	if err != nil {
		panic(err)
	}

	permissionsClient, err := clients.NewPermissionsClient(permissionsPath)

	chatService := chat.New(storage)

	grpcApp := grpcapp.New(log, chatService, grpcPort)

	return &App{
		GRPCServer: grpcApp,
	}
}
