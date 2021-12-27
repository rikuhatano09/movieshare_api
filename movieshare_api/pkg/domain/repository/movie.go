package repository

import "github.com/rikuhatano09/movieshare_api/pkg/domain/model"

type (
	MovieRepository interface {
		FindMovieAtRandom(*string) (model.Movie, error)
		CreateMovie(model.Movie) (model.Movie, error)
		GetMovieList(*string, *string, *string) ([]model.Movie, error)
		GetMovieByID(uint64) (model.Movie, error)
		PutMovie(model.Movie, uint64) (model.Movie, error)
	}
)
