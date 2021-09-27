package handlers

import (
	"MovieProject/entities"
	"MovieProject/repo"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
)

func NewServer() *http.Server {
	r := mux.NewRouter()
	//route request to a func to pass as a parameter to method
	r.HandleFunc("/entities", PostNewMovie).Methods("POST")
	//handle is a reference to an object or structure
	svr := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8080",
		//configure the server to IP address
	}

	return svr
}

func PostNewMovie(w http.ResponseWriter, r *http.Request) {
	mv := entities.Movie{}

	err := json.NewDecoder(r.Body).Decode(&mv)
	if err != nil {
		fmt.Println(err)
	}

	mv.GetId()

	movieDb := repo.Db{}
	movieDb.PutToDb(mv)

	movieBytes, err := json.MarshalIndent(movieDb, "", "	")
	if err != nil {
		fmt.Println(err)
	}

	err = ioutil.WriteFile("moviedb.json", movieBytes, 0644)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(movieBytes)
}
