package usecase

import (
	"github.com/rikuhatano09/movieshare_api/internal/infrastructure/persistence"
	"github.com/rikuhatano09/movieshare_api/pkg/domain/contract"
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

func GetMovieList(title *string) ([]model.Movie, error) {
	moviePersistence := persistence.NewMoviePersistence()

	movieList, error := moviePersistence.GetMovieList(title)

	if error != nil {
		return []model.Movie{}, error
	}
	return movieList, nil
}

func CreateMovie(requestBody contract.MoviePostRequestBody) (model.Movie, error) {
	moviePersistence := persistence.NewMoviePersistence()

	movie := model.Movie{
		UserID:         requestBody.UserID,
		Title:          requestBody.Title,
		Overview:       requestBody.Overview,
		Genre:          requestBody.Genre,
		YouTubeLinkURL: requestBody.YouTubeLinkURL,
		GrinningScore:  nil,
	}

	movie.YouTubeThumbnailURL = movie.GetYouTubeThumbnailURL()

	movie, error := moviePersistence.CreateMovie(movie)

	if error != nil {
		return model.Movie{}, error
	}
	return movie, nil
}
