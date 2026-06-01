package main

import (
	"Courses/pkg/repository"
	"fmt"
	"log"
)

const connStr = "postgres://postgres:maxzak2012@localhost:5432/courses"

func main() {
	db, err := repository.New(connStr)
	if err != nil {
		log.Fatal(err.Error())
	}

	book, err := db.GetBookByID(4)
	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Printf("%d: %s, price: %d, author id: %d, genre id: %d\n", book.ID, book.Name, book.Price, book.AuthorID, book.GenreID)
}
