package api

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (api *api) books(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		strID, ok := mux.Vars(r)["id"]
		if ok {
			id, err := strconv.Atoi(strID)
			if (err != nil){
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			book, err := api.db.GetBookByID(id)
			if 
		}
	}
}
