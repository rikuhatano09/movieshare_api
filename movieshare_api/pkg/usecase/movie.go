package usecase

import (
	"github.com/rikuhatano09/movieshare_api/internal/infrastructure/persistence"
	"github.com/rikuhatano09/movieshare_api/pkg/domain/contract"
	"github.com/rikuhatano09/movieshare_api/pkg/domain/model"
)

func GetMovieAtRandom(genre *string) (model.Movie, error) {
	moviePersistence := persistence.NewMoviePersistence()

	movie, error := moviePersistence.FindMovieAtRandom(genre)

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

func GetMovieByID(id uint64) (model.Movie, error) {
	moviePersistence := persistence.NewMoviePersistence()

	movie, error := moviePersistence.GetMovieByID(id)

	if error != nil {
		return model.Movie{}, error
	}
	return movie, nil
}

func CreateMovie(requestBody contract.MoviePostRequestBody) (model.Movie, error) {
	moviePersistence := persistence.NewMoviePersistence()

	movie := model.Movie{
		UserID:         requestBody.UserID,
		UserName:       requestBody.UserName,
		Title:          requestBody.Title,
		Overview:       requestBody.Overview,
		Genre:          requestBody.Genre,
		YouTubeTitleID: requestBody.YouTubeTitleID,
		GrinningScore:  nil,
	}

	movie, error := moviePersistence.CreateMovie(movie)

	if error != nil {
		return model.Movie{}, error
	}
	return movie, nil
}

func PutMovie(requestBody contract.MoviePutRequestBody, id uint64) (model.Movie, error) {
	moviePersistence := persistence.NewMoviePersistence()

	movie, error := moviePersistence.PutMovie(requestBody.GrinningScore, id)

	if error != nil {
		return model.Movie{}, error
	}
	return movie, nil

}
