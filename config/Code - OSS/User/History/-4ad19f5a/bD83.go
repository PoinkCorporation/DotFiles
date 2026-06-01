package postgresql

import (
	"chat/internal/domain/models"
)

func (repo *PGRepo) NewChat(chat models.Chat) (int64, error) {
	const op = "storage.postgresql.NewChat"
	var id int64

	err := repo.pool.QueryRow()
}
