package repository

import (
	"Courses/pkg/models"
	"context"
)

func (repo *PGRepo) GetBooks() ([]models.Book, error) {
	rows, err := repo.pool.Query(context.Background(), `
		SELECT id, name, author_id, genre_id, price
		FROM books;
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var data []models.Book
	for rows.Next() {
		var book models.Book
		err = rows.Scan(
			&book.ID,
			&book.Name,
			&book.AuthorID,
			&book.GenreID,
			&book.Price,
		)
		if err != nil {
			return nil, err
		}

		data = append(data, book)
	}

	return data, nil
}

func (repo *PGRepo) NewAuthor(book models.Author) error {
	_, err := repo.pool.Exec(context.Background(), `
		INSERT INTO author (id, author_name)
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
