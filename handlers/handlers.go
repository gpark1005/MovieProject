package handlers

import (
	"encoding/json"
	"github.com/gpark1005/MovieProject/entities"
	"github.com/gpark1005/MovieProject/service"
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

func (mh MovieHandler) PostMovieHandler(w http.ResponseWriter, r *http.Request) {
	mv := entities.Movie{}

	err := json.NewDecoder(r.Body).Decode(&mv)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	err = mh.Svc.CreateNewMovie(mv)
	if err != nil {
		switch err.Error() {
		case "invalid rating":
			http.Error(w, err.Error(), http.StatusBadRequest)
		case "movie does not exist":
			http.Error(w, err.Error(), http.StatusNotFound)

		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

}

func (mh MovieHandler) GetMovieHandler(w http.ResponseWriter, r *http.Request) {
	return
}
