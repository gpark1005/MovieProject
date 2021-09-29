package repo

import (
	"MovieProject/entities"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Db struct {
	Movies []entities.Movie `json:"Movies"`
}

func (data *Db) PutToDb(movie entities.Movie) {
	data.Movies = append(data.Movies, movie)
}

func WriteNewMovie(film entities.Movie) (Db, error) {
	outPut, err := ioutil.ReadFile("moviedb.json")
	if err != nil {
		fmt.Println(err)
	}

	film.GetId()

	dataBase := Db{}
	err = json.Unmarshal(outPut, &dataBase)
	if err != nil {
		fmt.Println(err)
	}

	dataBase.Movies = append(dataBase.Movies, film)

	input, err := json.MarshalIndent(dataBase, "", " ")
	if err != nil {
		fmt.Println(err)
	}
	err = ioutil.WriteFile("moviedb.json", input, 0664)
	return dataBase, err
}
