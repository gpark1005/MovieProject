package handlers

import (
	"MovieProject/entities"
	"MovieProject/service"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type MovieHandler struct {
	Svc service.Service
}

func NewMovieHandler(s service.Service) MovieHandler {
	return MovieHandler{
		Svc: s,
	}
}

func (movieOne MovieHandler) PostNewMovie(w http.ResponseWriter, r *http.Request) {
	mv := entities.Movie{}

	err := json.NewDecoder(r.Body).Decode(&mv)
	if err != nil {
		fmt.Println(err)
	}

	err = movieOne.Svc.AddMovie(mv)
	if err != nil {
		switch err.Error() {
		case "movie already exists":
			http.Error(w, err.Error(), http.StatusBadRequest)
		case "invalid rating":
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}

func (mh MovieHandler) GetMovieHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	id := params["Id"]
	fmt.Println(id)

}

func (mov MovieHandler) GetById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["Id"]

	movById, err := mov.Svc.FindById(id)
	if err != nil {
		switch err.Error() {
		case "movie not found":
			http.Error(w, err.Error(), http.StatusNotFound)
		}
	}

	movie, err := json.MarshalIndent(movById, "", "	")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	_, _ = w.Write(movie)
}

func (mov MovieHandler) DeleteMovie(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["Id"]

	err := mov.Svc.DeleteMovie(id)
	if err != nil {
		switch err.Error() {
		case "failed to delete movie - does not exist":
			http.Error(w, err.Error(), http.StatusNotFound)
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
