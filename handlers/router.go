package handlers

import (
	"github.com/gorilla/mux"
)

type Server struct {
}

func ConfigureRouter(handler MovieHandler) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/movie", handler.PostMovieHandler).Methods("POST")
	r.HandleFunc("/movie/{id}", handler.PostMovieHandler).Methods("POST")

	return r

}
