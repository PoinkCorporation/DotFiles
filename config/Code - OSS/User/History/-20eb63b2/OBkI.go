package repository

import (
	"Courses/pkg/models"
	"context"
)

func (repo *PGRepo) GetAuthors() ([]models.Author, error) {
	rows, err := repo.pool.Query(context.Background(), `
		SELECT id, author_name
		FROM author;
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var data []models.Author
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

func (repo *PGRepo) NewAuthor(book models.Author) error {
	_, err := repo.pool.Exec(context.Background(), `
		INSERT INTO author (author_name)
		VALUES ($1);
		`,
		book.Name,
	)

	return err
}

func (repo *PGRepo) GetAuthorByID(id int) (models.Author, error) {
	var author models.Author
	err := repo.pool.QueryRow(context.Background(), `
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
