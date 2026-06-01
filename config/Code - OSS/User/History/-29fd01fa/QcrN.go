package main

import (
	"Courses/pkg/repository"
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

}
