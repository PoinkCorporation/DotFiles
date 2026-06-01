package app

import grpcapp "chat/internal/app/grpc"

type App struct {
	GRPCServer *grpcapp.App
}
