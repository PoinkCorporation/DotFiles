package permissions

import (
	"context"
	permv1 "permissions/gen/go/permissions"
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
