package usecase

import (
	"github.com/rikuhatano09/movieshare_api/internal/infrastructure/persistence"
	"github.com/rikuhatano09/movieshare_api/pkg/domain/model"
)

func GetMovieAtRandom() (model.Movie, error) {
	moviePersistence := persistence.NewMoviePersistence()

	movie, error := moviePersistence.FindMovieAtRandom()

	if error != nil {
		return model.Movie{}, error
	}
	return movie, nil
}
