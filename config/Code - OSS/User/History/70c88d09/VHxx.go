package app

import (
	"log/slog"
	grpcapp "permissions/internal/app/grpc"
	"permissions/internal/services/permissions"
	"permissions/internal/storage/postgresql"
	"time"
)

type App struct {
	GRPCServer *grpcapp.App
}

func New(
	log *slog.Logger,
	grpcPort int,
	storagePath string,
	tokenTTL time.Duration,
) *App {
	storage, err := postgresql.New(storagePath)
	if err != nil {
		panic(err)
	}

	authService := permissions.New(log, storage)

	grpcApp := grpcapp.New(log, authService, grpcPort)

	return &App{
		GRPCServer: grpcApp,
	}
}
