package main

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v4"
)

type book struct {
	id        int
	name      string
	author_id int
	genre_id  int
	price     int
}

func main() {
	conn, err := pgx.Connect(context.Background(), "postgres://postgres:maxzak2012@localhost:5432/courses")
	if err != nil {
		log.Fatal(err.Error())
	}

	defer conn.Close(context.Background())

	var b book
	err = conn.QueryRow(
		context.Background(), "SELECT id, name, price FROM books WHERE id = 2;",
	).Scan(
		&b.id,
		&b.name,
		&b.price,
	)

	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println(b)
}
