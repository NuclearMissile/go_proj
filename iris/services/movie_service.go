package services

import (
	"fmt"
	"github.com/iris/repositories"
)

type MovieService interface {
	ShowMovieName() string
}

type MovieServiceManager struct {
	repo repositories.MovieRepository
}

func NewMovieServiceManager(repo repositories.MovieRepository) MovieService {
	return &MovieServiceManager{repo: repo}
}

func (m *MovieServiceManager) ShowMovieName() string {
	return fmt.Sprintf("Movie name: %s", m.repo.GetMovieName())
}
