package postgresql

import (
	"chat/internal/domain/models"
	"context"
)

func (repo *PGRepo) NewChat(ctx context.Context, chat models.Chat) (int64, error) {
	const op = "storage.postgresql.NewChat"
	var id int64

	err := repo.pool.QueryRow(
		ctx,
		`INSERT INTO chats(chat_type, title, created_by) VALUES($1, $2, $3) RETURNING id;`,
		chat.Chat_type,
		chat.Title,
		chat.Created_by
	)
}
