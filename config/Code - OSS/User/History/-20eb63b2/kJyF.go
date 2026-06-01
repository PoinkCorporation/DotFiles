package repository

import (
	"Courses/pkg/models"
	"context"
)

func (repo *PGRepo) GetAuthors() (data []models.Author, err error) {
	rows, err := repo.pool.Query(context.Background(), `
		SELECT id, author_name
		FROM author;
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var author models.Author
		err = rows.Scan(
			&author.ID,
			&author.Name,
		)
		if err != nil {
			return nil, err
		}

		data = append(data, author)
	}

	return data, nil
}

func (repo *PGRepo) NewAuthor(author models.Author) (err error) {
	_, err = repo.pool.Exec(context.Background(), `
		INSERT INTO author (author_name)
		VALUES ($1);
		`,
		author.Name,
	)

	return err
}

func (repo *PGRepo) GetAuthorByID(id int) (author models.Author, err error) {
	err = repo.pool.QueryRow(context.Background(), `
		SELECT id, author_name
		FROM author where id = $1;
		`,
		id,
	).Scan(
		&author.ID,
		&author.Name,
	)
	return author, err
}

func (repo *PGRepo) DeleteAuthorByID(id int) (err error) {
	_, err = repo.pool.Exec(
		context.Background(),
		`DELETE FROM author where id = $1;`,
		id,
	)
	return err
}

func (repo *PGRepo) UpdateAuthor(author models.Author) (err error) {
	_, err = repo.pool.Exec(
		context.Background(),
		`UPDATE books
		SET author_name = $1
		WHERE id = $2`,
		author.Name,
		author.ID,
	)

	return err
}
