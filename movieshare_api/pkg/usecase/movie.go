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

func GetMovieList(title *string, genre *string, userId *string) ([]model.Movie, error) {
	moviePersistence := persistence.NewMoviePersistence()

	movieList, error := moviePersistence.GetMovieList(title, genre, userId)

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

func PutMovie(requestBody contract.MoviePutRequestBody, id uint64) (model.Movie, error) {
	moviePersistence := persistence.NewMoviePersistence()

	movie := model.Movie{
		GrinningScore: requestBody.GrinningScore,
	}

	movie, error := moviePersistence.PutMovie(movie, id)

	if error != nil {
		return model.Movie{}, error
	}
	return movie, nil

}
