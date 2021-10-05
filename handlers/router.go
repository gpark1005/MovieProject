package handlers

import (
	"github.com/gorilla/mux"
)

func ConfigureRouter(handler MovieHandler) *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/movie", handler.PostNewMovie).Methods("POST")
	router.HandleFunc("/movie", handler.GetMovieHandler).Methods("GET")
	router.HandleFunc("/movie/{Id}", handler.GetById).Methods("GET")
	router.HandleFunc("/movie/{Id}", handler.DeleteMovie).Methods("DELETE")

	return router
}
