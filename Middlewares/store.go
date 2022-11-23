package Middlewares

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/varuns-zop/movie/Models"
)

type Store struct {
	db *sql.DB
}

func New(database *sql.DB) *Store {
	return &Store{db: database}
}

func (s *Store) CreateMovie(movie Models.Movie) (Models.Movie, error) {

	m := Models.Movie{}

	_, err := s.db.ExecContext(context.Background(), "CREATE TABLE IF NOT EXISTS Movies (id INT AUTO_INCREMENT PRIMARY KEY, name VARCHAR(255), genre VARCHAR(255), rating FLOAT, plot VARCHAR(255), released INT); ")
	if err != nil {
		return m, errors.New(err.Error())
	}

	result, insertErr := s.db.ExecContext(context.Background(), "INSERT INTO Movies(name, genre, rating, plot, released) VALUES (?,?,?,?,?)", movie.Name, movie.Genre, movie.Rating, movie.Plot, movie.Released)
	if insertErr != nil {
		return m, errors.New(insertErr.Error())
	}

	id, e := result.LastInsertId()
	if e != nil {
		return m, errors.New(e.Error())
	}

	data := s.db.QueryRowContext(context.Background(), "SELECT id,name,released,rating,plot,genre FROM Movies WHERE id = (?)", id)
	if errS := data.Scan(&m.Id, &m.Name, &m.Released, &m.Rating, &m.Plot, &m.Genre); errS != nil {
		fmt.Errorf("scan error %v", errS.Error())
	}
	return m, nil
}

func (s *Store) UpdateMovieByID(movie Models.Movie, params map[string]string) (Models.Movie, error) {

	m := Models.Movie{}

	_, errorS := s.db.QueryContext(context.Background(), "UPDATE  Movies SET released=?,rating=?,plot=?,genre=? WHERE id = (?)", movie.Released, movie.Rating, movie.Plot, movie.Genre, params["id"])
	if errorS != nil {
		return m, errors.New(errorS.Error())
	}

	data := s.db.QueryRowContext(context.Background(), "SELECT id,name,released,rating,plot,genre FROM Movies WHERE id = (?)", params["id"])

	if errS := data.Scan(&m.Id, &m.Name, &m.Released, &m.Rating, &m.Plot, &m.Genre); errS != nil {
		fmt.Errorf("scan error %v", errS.Error())
	}

	return m, nil
}

func (s *Store) DeleteMovieByID(params map[string]string) (Models.GenericResponse, error) {

	_, errorS := s.db.QueryContext(context.Background(), "DELETE FROM Movies WHERE id = (?)", params["id"])
	if errorS != nil {
		return Models.GenericResponse{Code: 404, Status: "FAILURE", Data: "ID Not Found."}, errors.New(errorS.Error())
	}
	return Models.GenericResponse{Code: 200, Status: "SUCCESS", Data: "Movie deleted successfully."}, nil
}

func (s *Store) GetALlMovies() ([]Models.Movie, error) {
	var movies []Models.Movie

	data, fetchErr := s.db.QueryContext(context.Background(), "SELECT id,name,released,rating,plot,genre FROM Movies")
	if fetchErr != nil {
		return nil, errors.New(fetchErr.Error())
	}

	m := Models.Movie{}
	for data.Next() {
		if errS := data.Scan(&m.Id, &m.Name, &m.Released, &m.Rating, &m.Plot, &m.Genre); errS != nil {
			fmt.Errorf("scan error %v", errS.Error())
		}
		movies = append(movies, m)
	}

	return movies, nil
}
