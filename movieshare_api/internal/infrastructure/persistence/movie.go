package persistence

import (
	"fmt"

	"github.com/jinzhu/gorm"

	"github.com/rikuhatano09/movieshare_api/pkg/domain/model"
	"github.com/rikuhatano09/movieshare_api/pkg/domain/repository"
)

type (
	// MoviePersistence is a persistence to preserve movies.
	MoviePersistence struct {
		Connection *gorm.DB
	}
)

// NewMoviePersistence creates a new movie persistence.
func NewMoviePersistence() repository.MovieRepository {
	return MoviePersistence{
		Connection: getDBConnection(),
	}
}

// FindMovieAtRandom find a movie at random.
func (moviePersistence MoviePersistence) FindMovieAtRandom(genre *string) (model.Movie, error) {
	movie := model.Movie{}

	query := moviePersistence.Connection.New().Table("movie")

	if genre != nil {
		query = query.Where(`"genre" = ?`, *genre)
	}

	result := query.Order("random()").Find(&movie)

	if result.RecordNotFound() {
		return movie, nil
	}
	if result.Error != nil {
		return movie, result.Error
	}
	return movie, nil
}

// GetMovieList get movies under the specified conditions.
func (moviePersistence MoviePersistence) GetMovieList(title *string, genre *string, userId *string) ([]model.Movie, error) {
	movieList := []model.Movie{}

	query := moviePersistence.Connection.New().Table("movie")

	if title != nil {
		query = query.Where(`"title" LIKE ?`, "%"+*title+"%")
	}
	if genre != nil {
		query = query.Where("genre = ?", *genre)
	}
	if userId != nil {
		query = query.Where("user_id = ?", *userId)
	}

	result := query.Find(&movieList)

	if result.RecordNotFound() {
		return movieList, nil
	}
	if result.Error != nil {
		return movieList, result.Error
	}
	return movieList, nil
}

// CreateMovie create a movie.
func (moviePersistence MoviePersistence) CreateMovie(movie model.Movie) (model.Movie, error) {
	result := moviePersistence.Connection.New().
		Table("movie").
		Create(&movie)

	if result.Error != nil {
		return movie, result.Error
	}
	return movie, nil
}

func (moviePersistence MoviePersistence) FindMovieByID(id uint64) (model.Movie, error) {
	movie := model.Movie{}

	result := moviePersistence.Connection.New().
		Table("movie").
		Where("id = ?", id).
		Find(&movie)

	if result.RecordNotFound() {
		return movie, nil
	}
	if result.Error != nil {
		return movie, result.Error
	}
	return movie, nil
}

func (moviePersistence MoviePersistence) UpdateMovie(attributes map[string]interface{}, id uint64) (model.Movie, error) {
	movie := model.Movie{ID: id}

	result := moviePersistence.Connection.New().
		Model(&movie).
		Updates(attributes)

	if result.RowsAffected == 0 {
		return movie, fmt.Errorf("record [ID: %d] does not exists", id)
	}
	if result.Error != nil {
		return movie, result.Error
	}
	// Find a movie by ID for sending a response.
	movie, error := moviePersistence.FindMovieByID(id)
	if error != nil {
		return movie, error
	}
	return movie, nil
}
