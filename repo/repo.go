package repo

import (
	"MovieProject/entities"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
)

type Db struct {
	Movies []entities.Movie `json:"Movies"`
}

type Repo struct {
	Filename string
}

func NewRepo(f string) Repo {
	return Repo{
		Filename: f,
	}
}

func (data *Db) PutToDb(movie entities.Movie) {
	data.Movies = append(data.Movies, movie)
}

func (r Repo) AddMovie(m entities.Movie) error {
	movieSlice := Db{}

	file, err := ioutil.ReadFile(r.Filename)
	if err != nil {
		return err
	}

	err = json.Unmarshal(file, &movieSlice)
	if err != nil {
		return err
	}

	for _, val := range movieSlice.Movies {
		if val.Title == m.Title {
			return errors.New("movie exits")
		}
	}

	movieSlice.Movies = append(movieSlice.Movies, m)

	movieBytes, err := json.MarshalIndent(movieSlice, "", "	")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(r.Filename, movieBytes, 0644)
	if err != nil {
		return err
	}

	return nil
}

func (r Repo) FindById(id string) (entities.Movie, error) {
	file, err := ioutil.ReadFile(r.Filename)
	if err != nil {
		fmt.Println(err)
	}

	movies := Db{}
	err = json.Unmarshal(file, &movies)

	match := entities.Movie{}

	for _, val := range movies.Movies {
		if val.Id == id {
			match = val
			return match, nil
		}
	}
	return entities.Movie{}, errors.New("movie not found")
}

func (r Repo) DeleteMovie(id string) error {
	file, err := ioutil.ReadFile(r.Filename)
	if err != nil {
		return err
	}

	movies := Db{}
	newDb := Db{}
	err = json.Unmarshal(file, &movies)

	dbSize := len(movies.Movies)
	for _, val := range movies.Movies {
		if val.Id == id {
			continue
		} else {
			newDb.Movies = append(newDb.Movies, val)
		}
	}

	if len(newDb.Movies) == dbSize {
		return errors.New("failed to delete movie - does not exist")
	}

	movieBytes, err := json.MarshalIndent(newDb, "", "	")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile("moviedb.json", movieBytes, 0644)
	if err != nil {
		return err
	}

	return nil
}
