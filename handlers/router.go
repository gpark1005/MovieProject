package handlers

import (
	"github.com/gorilla/mux"
)

type Server struct {
}

func ConfigureRouter(handler MovieHandler) *mux.Router {
// Some other random lines
	// Some other random lines
	// Some other random lines
	r := mux.NewRouter()

	r.HandleFunc("/movie", handler.PostMovieHandler).Methods("POST")
	r.HandleFunc("/movie/{id}", handler.GetMovieHandler).Methods("GET")

	return r

}
