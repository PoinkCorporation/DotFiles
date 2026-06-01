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
		rows.Scan(
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

func (repo *PGRepo) NewBook(book models.Book) error {
	_, err := repo.pool.Exec(context.Background(), `
		INSERT INTO books (name, author_id, genre_id, price)
		VALUES ($1, $2, $3, $4);
		`,
		book.Name,
		book.AuthorID,
		book.GenreID,
		book.Price,
	)

	return err
}

func (repo *PGRepo) GetBookByID(id int) (models.Book, error) {
	var book models.Book
	err := repo.pool.QueryRow(context.Background(), `
		SELECT id, name, author_id, genre_id, price
		FROM books where id = $1;
		`,
		id,
	).Scan(
		&book.ID,
		&book.Name,
		&book.AuthorID,
		&book.GenreID,
		&book.Price,
	)
	if err != nil {
		return models.Book{}, err
	}

	return book, nil
}
