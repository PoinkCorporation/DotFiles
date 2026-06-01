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
		var item models.Book
		rows.Scan(
			&item.ID,
			&item.Name,
			&item.AuthorID,
			&item.GenreID,
			&item.Price,
		)
		if err != nil {
			return nil, err
		}

		data = append(data, item)
	}

	return data, nil
}

func (repo *PGRepo) NewBook(item models.Book) error {
	_, err := repo.pool.Exec(context.Background(), `
		INSERT INTO books (name, author_id, genre_id, price)
		VALUES ($1, $2, $3, $4);
		`,
		item.Name,
		item.AuthorID,
		item.GenreID,
		item.Price,
	)

	return err
}
