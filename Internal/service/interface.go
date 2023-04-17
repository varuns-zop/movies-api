package service

import (
	"github.com/varuns-zop/movie/Internal/models"
)

type ServicesHandler interface {
	GetALlMovieService() ([]models.Movie, error)
	CreateMovieService(movie models.Movie) (models.Movie, error)
	UpdateMovieByIDService(movie models.Movie, params map[string]string) (models.Movie, error)
	DeleteMovieByIDService(params map[string]string) (models.GenericResponse, error)
}
