package handlers

import (
	"MovieProject/entities"
	"MovieProject/repo"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func PostNewMovie(w http.ResponseWriter, r *http.Request) {
	mv := entities.Movie{}

	err := json.NewDecoder(r.Body).Decode(&mv)
	if err != nil {
		fmt.Println(err)
	}

	mv.GetId()

	movieDb := repo.Db{}
	movieDb.PutToDb(mv)

	jsonBytes, err := json.MarshalIndent(movieDb, "", "	")
	if err != nil {
		fmt.Println(err)
	}

	err = ioutil.WriteFile("moviedb.json", jsonBytes, 0644)
	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(jsonBytes)

}
