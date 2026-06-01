package main

import (
	"Courses/pkg/models"
	"Courses/pkg/repository"
	"fmt"
	"log"
)

type book struct {
	id        int
	name      string
	author_id int
	genre_id  int
	price     int
}

const connStr = "postgres://postgres:maxzak2012@localhost:5432/courses"

func main() {
	db, err := repository.New(connStr)
	if err != nil {
		log.Fatal(err.Error())
	}

	book := models.Book{
		AuthorID: 1,
		GenreID:  5,
		Price:    150000,
		Name:     "Idiot",
	}

	err = db.NewBook(book)

	data, err := db.GetBooks()
	if err != nil {
		log.Fatal(err.Error())
	}

	for _, item := range data {
		fmt.Printf("%d: %s, price: %d, author id: %d, genre id: %d\n", item.ID, item.Name, item.Price, item.AuthorID, item.GenreID)
	}
}
