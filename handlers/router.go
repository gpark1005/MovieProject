package handlers

import (
	"github.com/gorilla/mux"
)

func ConfigureRouter(handler MovieHandler) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/movie", handler.PostNewMovie).Methods("POST")
	r.HandleFunc("/movie", handler.GetMovieHandler).Methods("GET")
	r.HandleFunc("/movie/{Id}", handler.GetById).Methods("GET")
	r.HandleFunc("/movie/{Id}", handler.DeleteMovie).Methods("DELETE")

	return r
}
