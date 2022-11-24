package movie

import (
	"github.com/varuns-zop/movie/Internal/models"
	"github.com/varuns-zop/movie/Internal/stores"
)

type Service struct {
	store stores.StoreHandler
}

func New(s stores.StoreHandler) *Service {
	return &Service{store: s}
}

func (s *Service) GetALlMovieService() ([]models.Movie, error) {
	m, err := s.store.GetALlMovies()
	if err != nil {
		return nil, err
	}
	return m, err
}

func (s *Service) CreateMovieService(movie models.Movie) (models.Movie, error) {
	m, err := s.store.CreateMovie(movie)
	if err != nil {
		return models.Movie{}, err
	}
	return m, err
}

func (s *Service) UpdateMovieByIDService(movie models.Movie, params map[string]string) (models.Movie, error) {
	m, err := s.store.UpdateMovieByID(movie, params)
	if err != nil {
		return models.Movie{}, err
	}
	return m, err
}

func (s *Service) DeleteMovieByIDService(params map[string]string) (models.GenericResponse, error) {
	m, err := s.store.DeleteMovieByID(params)
	if err != nil {
		return models.GenericResponse{}, err
	}
	return m, err
}
