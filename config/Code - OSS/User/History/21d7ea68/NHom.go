package postgresql

import (
	"context"
	"fmt"
)

func (repo *PGRepo) UserRoles(ctx context.Context, userID int64) ([]string, error) {
	const op = "storage.postgresql.UserRoles"
	rows, err := repo.pool.Query(ctx,
		`SELECT r.name FROM roles r
		 JOIN user_roles ur ON ur.role_id = r.id
		 WHERE ur.user_id = $1`, userID)
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
