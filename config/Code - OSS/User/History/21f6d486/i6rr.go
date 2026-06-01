package api

import "net/http"

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello"))
}
