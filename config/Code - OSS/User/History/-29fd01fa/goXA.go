package main

import (
	"Courses/pkg/api"
	"Courses/pkg/repository"
	"log"

	"github.com/gorilla/mux"
)

const connStr = "postgres://postgres:maxzak2012@localhost:5432/courses"

func main() {
	db, err := repository.New(connStr)
	if err != nil {
		log.Fatal(err.Error())
	}

	api.New(mux.NewRouter(), db)
	api.Handle()
}
