package repositories

import "github.com/iris/datamodels"

type MovieRepository interface {
	GetMovieName() string
}

type MovieManager struct {
}

func NewMovieManager() MovieRepository {
	return &MovieManager{}
}

func (m *MovieManager) GetMovieName() string {
	movie := &datamodels.Movie{Name: "Hello, World."}
	return movie.Name
}
