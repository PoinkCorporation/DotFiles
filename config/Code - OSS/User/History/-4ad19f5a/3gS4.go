package postgresql

import (
	"chat/internal/domain/models"
	"chat/internal/storage"
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgconn"
)

func (repo *PGRepo) NewChat(ctx context.Context, chat models.Chat) (int64, error) {
	const op = "storage.postgresql.NewChat"
	var id int64

	err := repo.pool.QueryRow(
		ctx,
		`INSERT INTO chats(chat_type, title, created_by) VALUES($1, $2, $3) RETURNING id;`,
		chat.Chat_type,
		chat.Title,
		chat.Created_by,
	).Scan(
		&id,
	)

	if err != nil {
		var PgErr *pgconn.PgError

		if errors.As(err, &PgErr) && PgErr.Code == "23505" {
			return 0, fmt.Errorf("%s: %w", op, storage.ErrExists)
		}

		return 0, fmt.Errorf("%s: %w", op, err)
	}
}
