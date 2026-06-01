package main

import (
	"Courses/pkg/api"
	"Courses/pkg/repository"
	"log"

	"github.com/gorilla/mux"
)

const connStr = "postgres://postgres:maxzak2012@localhost:5432/courses"

func main() {
	router := mux.NewRouter()

	db, err := repository.New(connStr)
	if err != nil {
		log.Fatal(err.Error())
	}

	api := api.New(router, db)

	api.Handle()
	err = api.ListenAndServe(":8080")
	if err != nil {
		log.Fatal(err.Error())
	}
}
