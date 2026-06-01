package main

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v4"
)

type book struct {
	id     int
	author string
	price  int
	name   string
}

func main() {
	conn, err := pgx.Connect(context.Background(), "postgres://postgres:maxzak2012@localhost:5432/courses")
	if err != nil {
		log.Fatal(err.Error())
	}

	defer conn.Close(context.Background())

	conn.QueryRow(context.Background(), "CREATE TABLE books(id serial integer, author string, name string, price integer not null default 100);")
	row := conn.QueryRow(context.Background(), "SELECT * FROM books;")

	fmt.Print(row)
}
