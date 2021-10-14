package handlers

import (
	"github.com/gorilla/mux"
)

type Server struct {
}

func ConfigureRouter(handler MovieHandler) *mux.Router {
	//kjhkjhkj
	//kjhkjhkj
	//kjhkjhkj
	//kjhkjhkj
	r := mux.NewRouter()
	r.HandleFunc("/movie", handler.PostMovieHandler).Methods("POST")
	r.HandleFunc("/movie/{id}", handler.GetMovieHandler).Methods("GET")

	return r

}
