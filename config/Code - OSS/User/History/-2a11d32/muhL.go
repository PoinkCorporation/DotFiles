package postgresql

import (
	"context"
	"errors"
	"fmt"
	"sso/internal/domain/models"
	"sso/internal/storage"

	"github.com/jackc/pgx/v4"
)

func (repo *PGRepo) App(ctx context.Context, id int) (models.App, error) {
	const op = "storage.postgresql.App"

	var user models.User

	err := repo.pool.QueryRow(
		ctx,
		"SELECT id, email, pass_hash FROM users WHERE email = $1",
		email,
	).Scan(
		&user.ID,
		&user.Email,
		&user.PassHash,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return models.App{}, fmt.Errorf("%s: %w", op, storage.ErrAppNotFound)
		}

		return models.App{}, fmt.Errorf("%s: %w", op, err)
	}
}
