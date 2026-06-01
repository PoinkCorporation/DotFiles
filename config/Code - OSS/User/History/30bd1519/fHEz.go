package suite

import (
	chatv1 "chat/gen/go"
	"chat/internal/config"
	"context"
	"net"
	"strconv"
	"testing"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Suite struct {
	*testing.T                          // Потребуется для вызова методов *testing.T
	Cfg        *config.Config           // Конфигурация приложения
	ChatClient chatv1.ChatServiceClient // Клиент для взаимодействия с gRPC-сервером Auth
}

func New(t *testing.T) (context.Context, *Suite) {
	t.Helper()
	t.Parallel()

	cfg := config.MustLoad()

	ctx, cancelCtx := context.WithTimeout(context.Background(), cfg.GRPC.Timeout)

	t.Cleanup(func() {
		t.Helper()
		cancelCtx()
	})

	grpcAddress := net.JoinHostPort("localhost", strconv.Itoa(cfg.GRPC.Port))

	cc, err := grpc.DialContext(context.Background(),
		grpcAddress,
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		t.Fatalf("grpc server connection failed: %v", err)
	}

	chatClient := chatv1.NewChatServiceClient(cc)

	return ctx, &Suite{
		T:          t,
		Cfg:        cfg,
		ChatClient: chatClient,
	}
}
