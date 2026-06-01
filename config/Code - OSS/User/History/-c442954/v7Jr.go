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

func (repo *PGRepo) NewBook(book models.Book) (err error) {
	_, err = repo.pool.Exec(context.Background(), `
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

func (repo *PGRepo) GetBookByID(id int) (book models.Book, err error) {
	err = repo.pool.QueryRow(context.Background(), `
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
	return book, err
}

func (repo *PGRepo) DeleteBookByID(id int) (err error) {
	_, err = repo.pool.Exec(
		context.Background(),
		`DELETE FROM books where id = $1;`,
		id,
	)
	return err
}

func (repo *PGRepo) UpdateBookByID(id, authorID, genreID, price int, name string) (err error) {
	_, err = repo.pool.Exec(
		context.Background(),
		`UPDATE books
		SET author_id = $1, genre_id = $2, price = $3, name = $4
		WHERE id = $5`,
		authorID,
		genreID,
		price,
		name,
		id,
	)

	return err
}
