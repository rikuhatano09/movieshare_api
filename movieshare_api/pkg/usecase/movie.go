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

func GetMovieList(title *string, genre *string, user_id *string) ([]model.Movie, error) {
	moviePersistence := persistence.NewMoviePersistence()

	movieList, error := moviePersistence.GetMovieList(title, genre, user_id)

	if error != nil {
		return []model.Movie{}, error
	}
	return movieList, nil
}

func GetMovieID(id uint64) (model.Movie, error) {
	moviePersistence := persistence.NewMoviePersistence()

	movie, error := moviePersistence.GetMovieID(id)

	if error != nil {
		return model.Movie{}, error
	}
	return movie, nil
}
