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
