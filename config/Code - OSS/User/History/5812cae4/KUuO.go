package postgresql

import (
	"context"
	"errors"
	"fmt"
	"sso/internal/storage"

	"github.com/jackc/pgconn"
)

func (repo *PGRepo) SaveUser(ctx context.Context, email string, passHash []byte) (int64, error) {
	const op = "storage.postgresql.SaveUser"
	var id int64

	err := repo.pool.QueryRow(
		ctx,
		`INSERT INTO users(email, pass_hash) VALUES($1, $2) RETURNING id;`,
		email,
		passHash,
	).Scan(&id)
	if err != nil {
		// 1. Объявляем переменную для ошибки pgx
		var pgErr *pgconn.PgError

		// 2. Проверяем, является ли ошибка нарушением уникальности (Unique Violation)
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return 0, fmt.Errorf("%s: %w", op, storage.ErrUserExists)
		}

		// Любая другая ошибка БД
		return 0, fmt.Errorf("%s: %w", op, err)
	}

	return id, nil
}
