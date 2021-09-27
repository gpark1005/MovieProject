package repo

import (
	"MovieProject/entities"
)

type Db struct {
	Movies []entities.Movie `json:"Movies"`
}

func (data *Db) PutToDb(movie entities.Movie) {
	data.Movies = append(data.Movies, movie)
}
