package service

import (
	"errors"
	"github.com/google/uuid"
	"github.com/gpark1005/MovieProject/entities"
	"github.com/gpark1005/MovieProject/repository"
)

type Service struct {
	Repo repository.Repo
}

func NewService(r repository.Repo) Service {
	return Service{
		Repo: r,
	}
}

func (s Service) CreateNewMovie(mv entities.Movie) error {

	mv.Id = uuid.New().String()

	if mv.Rating <= 10 && mv.Rating >= 0 {
		err := s.Repo.CreateNewMovie(mv)
		if err != nil {
			return err
		}
		return nil
	}
	return errors.New("invalid Rating")
}
