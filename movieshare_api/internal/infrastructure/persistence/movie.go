package persistence

import (
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
func (moviePersistence MoviePersistence) FindMovieAtRandom() (model.Movie, error) {
	movie := model.Movie{}

	result := moviePersistence.Connection.New().
		Table("movie").
		Order("random()").
		Find(&movie)

	if result.RecordNotFound() {
		return movie, nil
	}
	if result.Error != nil {
		return movie, result.Error
	}
	return movie, nil
}

// GetMovieList get movies under the specified conditions.
func (moviePersistence MoviePersistence) GetMovieList(title *string, genre *string, user_id *string) ([]model.Movie, error) {
	movieList := []model.Movie{}

	query := moviePersistence.Connection.New().Table("movie")

	if title != nil {
		query = query.Where(`"title" LIKE ?`, "%"+*title+"%")
	}
	if genre != nil {
		query = query.Where("genre = ?", *genre)
	}
	if user_id != nil {
		query = query.Where("user_id = ?", *user_id)
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

func (moviePersistence MoviePersistence) GetMovieID(id uint64) (model.Movie, error) {
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
