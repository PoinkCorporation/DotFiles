package permissions

import (
	"context"
	"errors"
	permv1 "permissions/gen/go/permissions"
	"permissions/internal/storage"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Permissions interface {
	HasPermission(ctx context.Context, userID int64, permission string) (bool, error)
	UserRoles(ctx context.Context, userID int64) ([]string, error)
	AssignRole(ctx context.Context, userID int64, roleName string) error
}

type serverAPI struct {
	permv1.UnimplementedPermissionsServer
	permissions Permissions
}

func Register(gRPCServer *grpc.Server, permissions Permissions) {
	permv1.RegisterPermissionsServer(gRPCServer, &serverAPI{permissions: permissions})
}

func (s *serverAPI) HasPermission(ctx context.Context, req *permv1.HasPermissionRequest) (*permv1.HasPermissionResponse, error) {
	if req.GetUserId() == 0 {
		return nil, status.Error(codes.InvalidArgument, "user_id is required")
	}
	if req.GetRole() == "" {
		return nil, status.Error(codes.InvalidArgument, "role is required")
	}

	ok, err := s.permissions.HasPermission(ctx, req.GetUserId(), req.GetRole())

	if err != nil {
		if errors.Is(err, storage.ErrPermissionDenied) {
			return nil, status.Error(codes.PermissionDenied, "permission denied")
		}

		return nil, status.Error(codes.Internal, "failed to check permission")
	}

	return &permv1.HasPermissionResponse{Exists: ok}, nil
}

func (s *serverAPI) UserRoles(ctx context.Context, req *permv1.UserRolesRequest) (*permv1.UserRolesResponse, error) {

}

func (s *serverAPI) AssignRole(ctx context.Context, req *permv1.AssignRoleRequest) (*emptypb.Empty, error) {

}
