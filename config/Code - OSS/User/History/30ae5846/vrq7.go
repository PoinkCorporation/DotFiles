package repository

import (
	"Courses/pkg/models"
	"context"
)

func (repo *PGRepo) GetGenres() (data []models.Genre, err error) {
	rows, err := repo.pool.Query(context.Background(), `
		SELECT id, genre
		FROM books;
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var genre models.Genre
		err = rows.Scan(
			&genre.ID,
			&genre.Genre,
		)
		if err != nil {
			return nil, err
		}

		data = append(data, genre)
	}

	return data, nil
}

func (repo *PGRepo) NewGenre(genre models.Genre) (err error) {
	_, err = repo.pool.Exec(context.Background(), `
		INSERT INTO genre (genre)
		VALUES ($1);
		`,
		genre.Genre,
	)

	return err
}

func (repo *PGRepo) GetGenreByID(id int) (genre models.Genre, err error) {
	err = repo.pool.QueryRow(context.Background(), `
		SELECT id, genre
		FROM genre where id = $1;
		`,
		id,
	).Scan(
		&genre.ID,
		&genre.Genre,
	)
	return genre, err
}

func (repo *PGRepo) DeleteGenreByID(id int) (err error) {
	_, err = repo.pool.Exec(
		context.Background(),
		`DELETE FROM genre where id = $1;`,
		id,
	)
	return err
}

func (repo *PGRepo) UpdateGenreByID(id int, genre string) (err error) {
	_, err = repo.pool.Exec(
		context.Background(),
		`UPDATE genre
		SET genre = $1
		WHERE id = $2`,
		genre,
		id,
	)

	return err
}
