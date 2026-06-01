package api

import (
	"Courses/pkg/repository"
	"net/http"

	"github.com/gorilla/mux"
)

type api struct {
	r  *mux.Router
	db *repository.PGRepo
}

func New(router *mux.Router, db *repository.PGRepo) *api {
	return &api{r: router, db: db}
}

func (api *api) Handle() {
	api.r.HandleFunc("/api/books", api.books).Queries("id", "{id}")
	api.r.HandleFunc("/api/books", api.books)
	api.r.HandleFunc("/api/authors", api.authors).Queries("id", "{id}")
	api.r.HandleFunc("/api/authors", api.authors)
	api.r.HandleFunc("/api/genres", api.genres).Queries("id", "{id}")
	api.r.HandleFunc("/api/genres", api.genres)
}

func (api *api) ListenAndServe(addr string) error {
	return http.ListenAndServe(addr, api.r)
}
