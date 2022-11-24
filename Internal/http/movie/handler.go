package movie

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/varuns-zop/movie/Internal/models"
	"github.com/varuns-zop/movie/Internal/service"
	"github.com/varuns-zop/movie/Middlewares"
)

type serviceConnector struct {
	service service.ServicesHandler
}

func New(service service.ServicesHandler) *serviceConnector {
	return &serviceConnector{service: service}
}

func (c *serviceConnector) GetALLMovies(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	data, err := c.service.GetALlMovieService()
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		err := json.NewEncoder(w).Encode(models.GenericResponse{Code: 404, Status: "FAILURE", Data: "no data found"})
		Middlewares.CheckNillError(err)
		return
	}

	err = json.NewEncoder(w).Encode(data)
	Middlewares.CheckNillError(err)
	return
}

func (c *serviceConnector) AddingMovie(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	// Checking if request body is nil or not
	if r.Body == nil {
		err := json.NewEncoder(w).Encode(models.GenericResponse{Code: 404, Status: "FAILURE", Data: "no request body found"})
		Middlewares.CheckNillError(err)
		return
	}

	//Decoding the request data with struct type movie and checking if it is nil or not
	var movie models.Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)

	if movie.IsEmpty() {
		err := json.NewEncoder(w).Encode(models.GenericResponse{Code: 404, Status: "FAILURE", Data: "no data found in json"})
		Middlewares.CheckNillError(err)
		return
	}

	// creating the random value for movie id to be randomly created at a time

	data, err := c.service.CreateMovieService(movie)
	if err != nil {
		json.NewEncoder(w).Encode(models.GenericResponse{Code: 404, Status: "FAILURE", Data: "no request body found"})
	}
	response := models.MovieDetails{Code: 200, Status: "SUCCESS", Data: &data}
	err = json.NewEncoder(w).Encode(response)
	Middlewares.CheckNillError(err)
	return
}

func (c *serviceConnector) EditMovieById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Checking if request body is nil or not
	if r.Body == nil {
		err := json.NewEncoder(w).Encode(models.GenericResponse{Code: 404, Status: "FAILURE", Data: "no request body found"})
		Middlewares.CheckNillError(err)
		return
	}

	//Decoding the request data with struct type movie and checking if it is nil or not
	var movie models.Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)

	if movie.IsEmpty() {
		err := json.NewEncoder(w).Encode(models.GenericResponse{Code: 404, Status: "FAILURE", Data: "no data found in json"})
		Middlewares.CheckNillError(err)
		return
	}

	params := mux.Vars(r)

	data, err := c.service.UpdateMovieByIDService(movie, params)
	response := models.MovieDetails{Code: 200, Status: "SUCCESS", Data: &data}
	if err != nil {
		err := json.NewEncoder(w).Encode(models.GenericResponse{Code: 404, Status: "FAILURE", Data: "Invalid"})
		Middlewares.CheckNillError(err)
		return
	}

	err = json.NewEncoder(w).Encode(response)
	Middlewares.CheckNillError(err)
	return
}

func (c *serviceConnector) DeleteMovieById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	data, errorS := c.service.DeleteMovieByIDService(params)
	if errorS != nil {
		err := json.NewEncoder(w).Encode(models.GenericResponse{Code: 404, Status: "FAILURE", Data: "Id Not Found"})
		Middlewares.CheckNillError(err)
		return
	}

	err := json.NewEncoder(w).Encode(data)
	Middlewares.CheckNillError(err)
	return
}
