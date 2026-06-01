package postgresql

import (
	"context"
	"errors"
	"fmt"
	"sso/internal/storage"

	"github.com/jackc/pgx/v4"
)

func (repo *PGRepo) UserRoles(ctx context.Context, userID int64) ([]string, error) {
	const op = "storage.postgresql.UserRoles"

	rows, err := repo.pool.Query(
		ctx,
		`SELECT r.name FROM roles r
		 JOIN user_roles ur ON ur.role_id = r.id
		 WHERE ur.user_id = $1`,
		userID,
	)

	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	defer rows.Close()

	var roles []string

	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}
		roles = append(roles, name)
	}

	return roles, rows.Err()
}

func (repo *PGRepo) HasPermission(ctx context.Context, userID int64, permission string) (bool, error) {
	const op = "storage.postgresql.HasPermission"
	var exists bool
	err := repo.pool.QueryRow(ctx,
		`SELECT EXISTS (
			SELECT 1 FROM user_roles ur
			JOIN role_permissions rp ON rp.role_id = ur.role_id
			JOIN permissions p ON p.id = rp.permission_id
			WHERE ur.user_id = $1 AND p.name = $2
		)`,
		userID,
		permission,
	).Scan(&exists)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return false, fmt.Errorf("%s: %w", op, storage.ErrPermissionDenied)
		}

		return false, fmt.Errorf("%s: %w", op, err)
	}

	return exists, nil
}
