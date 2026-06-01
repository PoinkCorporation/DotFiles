package postgresql

import (
	"chat/internal/domain/models"
	"chat/internal/storage"
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
)

func (repo *PGRepo) NewChat(ctx context.Context, chat models.Chat) (models.Chat, error) {
	const op = "storage.postgresql.NewChat"

	err := repo.pool.QueryRow(
		ctx,
		`INSERT INTO chats(chat_type, title, created_by) VALUES($1, $2, $3) RETURNING id, chat_type, title, created_by;`,
		chat.ChatType,
		chat.Title,
		chat.CreatedBy,
	).Scan(
		&chat.ID,
		&chat.ChatType,
		&chat.Title,
		&chat.CreatedBy,
	)

	if err != nil {
		var PgErr *pgconn.PgError

		if errors.As(err, &PgErr) && PgErr.Code == "23505" {
			return models.Chat{}, fmt.Errorf("%s: %w", op, storage.ErrExists)
		}

		return models.Chat{}, fmt.Errorf("%s: %w", op, err)
	}

	return chat, nil
}

func (repo *PGRepo) ChatByID(ctx context.Context, id int64) (models.Chat, error) {
	const op = "storage.postgresql.ChatByID"
	var chat models.Chat

	err := repo.pool.QueryRow(
		ctx,
		`SELECT id, chat_type, title, created_by FROM chats WHERE id = $1;`,
		id,
	).Scan(
		&chat.ID,
		&chat.ChatType,
		&chat.Title,
		&chat.CreatedBy,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return models.Chat{}, fmt.Errorf("%s: %w", op, storage.ErrNotFound)
		}
		return models.Chat{}, fmt.Errorf("%s: %w", op, err)
	}

	return chat, nil
}

func (repo *PGRepo) UpdateChat(ctx context.Context, chat models.Chat) error {
	const op = "storage.postgresql.UpdateChat"

	_, err := repo.pool.Exec(
		ctx,
		`UPDATE chats SET chat_type = $1, title = $2 WHERE id = $3;`,
		chat.ChatType,
		chat.Title,
		chat.ID,
	)

	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (repo *PGRepo) DeleteChat(ctx context.Context, id int64) error {
	const op = "storage.postgresql.DeleteChat"

	_, err := repo.pool.Exec(
		ctx,
		`DELETE FROM chats WHERE id = $1;`,
		id,
	)

	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (repo *PGRepo) ListUserChats(ctx context.Context, userID int64) ([]models.Chat, error) {
	const op = "storage.postgresql.ListUserChats"

	rows, err := repo.pool.Query(
		ctx,
		`SELECT c.id, c.chat_type, c.title, c.created_by
		FROM chats c
		JOIN chat_users cu ON c.id = cu.chat_id
		WHERE cu.user_id = $1;`,
		userID,
	)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer rows.Close()

	var chats []models.Chat
	for rows.Next() {
		var chat models.Chat
		if err := rows.Scan(&chat.ID, &chat.ChatType, &chat.Title, &chat.CreatedBy); err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}
		chats = append(chats, chat)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return chats, nil
}

func (repo *PGRepo) AddChatUser(ctx context.Context, userID, chatID int64, role string) error {
	const op = "storage.postgresql.AddChatUser"

	_, err := repo.pool.Exec(
		ctx,
		`INSERT INTO chat_users(user_id, chat_id, user_role) VALUES($1, $2, $3);`,
		userID,
		chatID,
		role,
	)

	if err != nil {
		var PgErr *pgconn.PgError
		if errors.As(err, &PgErr) && PgErr.Code == "23505" {
			return fmt.Errorf("%s: %w", op, storage.ErrExists)
		}
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (repo *PGRepo) DeleteChatUser(ctx context.Context, userID, chatID int64) error {
	const op = "storage.postgresql.DeleteChatUser"

	_, err := repo.pool.Exec(
		ctx,
		`DELETE FROM chat_users WHERE user_id = $1 AND chat_id = $2;`,
		userID,
		chatID,
	)

	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	return nil
}

func (repo *PGRepo) ChatUsers(ctx context.Context, chatID int64) ([]models.ChatUser, error) {
	const op = "storage.postgresql.ChatUsers"

	rows, err := repo.pool.Query(
		ctx,
		`SELECT user_id, chat_id, user_role, joined_at FROM chat_users WHERE chat_id = $1;`,
		chatID,
	)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}
	defer rows.Close()

	var users []models.ChatUser
	for rows.Next() {
		var user models.ChatUser
		if err := rows.Scan(&user.UserID, &user.ChatID, &user.Role, &user.JoinedAt); err != nil {
			return nil, fmt.Errorf("%s: %w", op, err)
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return users, nil
}
