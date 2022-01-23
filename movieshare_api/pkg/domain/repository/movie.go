package repository

import "github.com/rikuhatano09/movieshare_api/pkg/domain/model"

type (
	MovieRepository interface {
		FindMovieAtRandom(*string) (model.Movie, error)
		CreateMovie(model.Movie) (model.Movie, error)
		GetMovieList(*string, *string, *string) ([]model.Movie, error)
		FindMovieByID(uint64) (model.Movie, error)
		UpdateMovie(*uint32, uint64) (model.Movie, error)
	}
)
