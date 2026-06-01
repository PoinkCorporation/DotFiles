package permissions

import (
	"context"
	"log/slog"
)

type PermissionsChecker interface {
	HasPermission(ctx context.Context, userID int64, permission string) (bool, error)
	UserRoles(ctx context.Context, userID int64) ([]string, error)
	AssignRole(ctx context.Context, userID int64, roleName string) error
}

type Permissions struct {
	log         *slog.Logger
	permChecker PermissionsChecker
}

func New(log *slog.Logger, permChecker PermissionsChecker) *Permissions {
	return &Permissions{log: log, permChecker: permChecker}
}

func (perm *Permissions) HasPermission(ctx context.Context, userID int64, permission string) (bool, error) {
	const op = "Permissions.HasPermission"

	log := perm.log.With(
		slog.String("op", op),
		slog.String("permission", permission),
		slog.Int64("userID", userID),
	)

	log.Info("Check permission for user: ", userID)

	exists, err := perm.permChecker. (ctx, userID, permission)
}
