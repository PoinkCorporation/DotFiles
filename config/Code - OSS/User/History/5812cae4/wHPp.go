package postgresql

import (
	"context"
	"fmt"
)

func (repo *PGRepo) SaveUser(ctx context.Context, email string, passHash []byte) (int64, error) {
	const op = "storage.sqlite.SaveUser"

	_, err := repo.pool.Exec(ctx, `INSERT INTO users(email, pass_hash) VALUES($1, $2);`, email, passHash)
	if err != nil {
		return 0, fmt.Errorf("%s: %w", op, err)
	}
}
