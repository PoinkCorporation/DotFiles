package postgresql

func (repo *PGRepo) UserRoles(ctx, userID int64) ([]string, error)
