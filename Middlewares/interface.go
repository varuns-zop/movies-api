package Middlewares

import "github.com/varuns-zop/movie/Models"

type StoreHandler interface {
	GetALlMovies() ([]Models.Movie, error)
	CreateMovie(movie Models.Movie) (Models.Movie, error)
	UpdateMovieByID(movie Models.Movie, params map[string]string) (Models.Movie, error)
	DeleteMovieByID(params map[string]string) (Models.GenericResponse, error)
}
