package stores

import (
	"github.com/varuns-zop/movie/Internal/models"
)

type StoreHandler interface {
	GetALlMovies() ([]models.Movie, error)
	CreateMovie(movie models.Movie) (models.Movie, error)
	UpdateMovieByID(movie models.Movie, params map[string]string) (models.Movie, error)
	DeleteMovieByID(params map[string]string) (models.GenericResponse, error)
}
