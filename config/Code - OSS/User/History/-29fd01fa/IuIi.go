package main

import (
	"context"
	"log"

	"github.com/jackc/pgx/v4"
)

type book struct {
	id        int
	author_id int
	genre_id  int
	price     int
	name      string
}

func main() {
	conn, err := pgx.Connect(context.Background(), "postgres://postgres:maxzak2012@localhost:5432/courses")
	if err != nil {
		log.Fatal(err.Error())
	}

	defer conn.Close(context.Background())

	var b book
	err = conn.QueryRow(
		context.Background(), "SELECT * FROM books;",
	).Scan(
		&b.id,
		&b.name,
		&b.author_id,
		&b.genre_id,
		&b.price,
	)

	if err != nil {
		log.Fatal(err.Error())
	}
}
