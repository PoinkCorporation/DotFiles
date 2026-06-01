package app

import (
	"log/slog"
	"time"

	grpcapp "sso/internal/app/grpc"
	clients "sso/internal/clients/permissions"
	"sso/internal/services/auth"
	"sso/internal/storage/postgresql"
)

type App struct {
	GRPCServer *grpcapp.App
}

func New(
	log *slog.Logger,
	grpcPort int,
	storagePath string,
	permissionsAddress string,
	tokenTTL time.Duration,
) *App {
	storage, err := postgresql.New(storagePath)
	if err != nil {
		panic(err)
	}

	permClient, err := clients.NewPermissionsClient(permissionsAddress)
	if err != nil {
		panic(err)
	}

	authService := auth.New(log, storage, storage, storage, permClient, tokenTTL)

	grpcApp := grpcapp.New(log, authService, grpcPort)

	return &App{
		GRPCServer: grpcApp,
	}
}
