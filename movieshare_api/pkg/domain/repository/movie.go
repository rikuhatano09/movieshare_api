package repository

import "github.com/rikuhatano09/movieshare_api/pkg/domain/model"

type (
	MovieRepository interface {
		FindMovieAtRandom() (model.Movie, error)
		GetMovieList(*string, *string, *string) ([]model.Movie, error)
		GetMovieID(uint64) (model.Movie, error)
	}
)
