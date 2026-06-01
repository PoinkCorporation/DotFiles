package postgresql

import "context"

func (repo *PGRepo) SaveUser(ctx context.Context, email string, passHash []byte) (int64, error) {
	row := repo.pool.QueryRow(ctx, `INSERT INTO users(email, pass_hash) VALUES($1, $2);`, email, passHash)
}
