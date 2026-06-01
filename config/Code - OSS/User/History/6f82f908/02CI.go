package permissions

import (
	"context"
	"log/slog"
	"permissions/internal/storage/postgresql"
)

type Permissions struct {
	log         *slog.Logger
	storage 	*postgresql.PGRepo
}

func New(log *slog.Logger, permChecker *postgresql.PGRepo) *Permissions {
	return &Permissions{log: log, }
}

func (perm *Permissions) HasPermission(ctx context.Context, userID int64, permission string) (bool, error) {
	const op = "Permissions.HasPermission"

	log := perm.log.With(
		slog.String("op", op),
		slog.String("permission", permission),
		slog.Int64("userID", userID),
	)

	log.Info("Check permission for user: ", userID)

	exists, err := 
}
