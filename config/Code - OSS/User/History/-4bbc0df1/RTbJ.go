package suite

import (
	"context"
	"net"
	permv1 "permissions/gen/go/permissions"
	"permissions/internal/config"
	"strconv"
	"testing"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Suite struct {
	*testing.T
	Cfg               *config.Config
	PermissionsClient permv1.PermissionsClient
}

func New(t *testing.T) (context.Context, *Suite) {
	t.Helper()
	t.Parallel()

	cfg := config.MustLoad()

	ctx, cancelCtx := context.WithTimeout(context.Background(), cfg.GRPC.Timeout)

	grpcAddress := net.JoinHostPort("localhost", strconv.Itoa(cfg.GRPC.Port))

	cc, err := grpc.NewClient(
		grpcAddress,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		t.Fatalf("grpc server connection failed: %v", err)
	}

	t.Cleanup(func() {
		t.Helper()
		cc.Close()
		cancelCtx()
	})

	permClient := permv1.NewPermissionsClient(cc)

	return ctx, &Suite{
		T:                 t,
		Cfg:               cfg,
		PermissionsClient: permClient,
	}
}
