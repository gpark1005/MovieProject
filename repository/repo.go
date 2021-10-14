package repository

import (
	"encoding/json"
	"github.com/gpark1005/MovieProject/entities"
	"io/ioutil"
)

type movieStruct struct {
	Movies []entities.Movie
}

type Repo struct {
	Filename string
}

func NewRepository(filename string) Repo {
	return Repo{
		Filename: filename,
	}
}

func (r Repo) CreateNewMovie(mv entities.Movie) error {

	ms := movieStruct{}

	bs, err := ioutil.ReadFile(r.Filename)
	if err != nil {
		return err
	}

	err = json.Unmarshal(bs, &ms)
	if err != nil {
		return err
	}

	ms.Movies = append(ms.Movies, mv)

	bytesSlice, err := json.MarshalIndent(ms, "", "	")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(r.Filename, bytesSlice, 0644)
	if err != nil {
		return err
	}

	return nil

}
