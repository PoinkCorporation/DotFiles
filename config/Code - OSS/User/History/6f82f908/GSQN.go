package permissions

import (
	"context"
	"log/slog"
)

type PermissionChecker interface {
	HasPermission(ctx context.Context, userID int64, permission string) (bool, error)
	UserRoles(ctx context.Context, userID int64) ([]string, error)
	AssignRole(ctx context.Context, userID int64, roleName string) error
}

type Permissions struct {
	log         *slog.Logger
	permChecker *PermissionChecker
}

func New(log *slog.Logger, permChecker *PermissionChecker) *Permissions {
	return &Permissions{log: log, permChecker: permChecker}
}

func (perm *Permissions) HasPermission(ctx context.Context, userID int64, permission string) (bool, error) {

}
