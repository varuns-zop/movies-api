package Controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/varuns-zop/movie/Middlewares"
	"github.com/varuns-zop/movie/Models"
)

type storeConnector struct {
	store Middlewares.StoreHandler
}

func Connect(s Middlewares.StoreHandler) *storeConnector {
	return &storeConnector{store: s}
}

func (c *storeConnector) GetALLMovies(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	data, err := c.store.GetALlMovies()
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		err := json.NewEncoder(w).Encode(Models.GenericResponse{Code: 404, Status: "FAILURE", Data: "no data found"})
		Middlewares.CheckNillError(err)
		return
	}

	err = json.NewEncoder(w).Encode(data)
	Middlewares.CheckNillError(err)
	return
}

func (c *storeConnector) AddingMovie(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	// Checking if request body is nil or not
	if r.Body == nil {
		err := json.NewEncoder(w).Encode(Models.GenericResponse{Code: 404, Status: "FAILURE", Data: "no request body found"})
		Middlewares.CheckNillError(err)
		return
	}

	//Decoding the request data with struct type movie and checking if it is nil or not
	var movie Models.Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)

	if movie.IsEmpty() {
		err := json.NewEncoder(w).Encode(Models.GenericResponse{Code: 404, Status: "FAILURE", Data: "no data found in json"})
		Middlewares.CheckNillError(err)
		return
	}

	// creating the random value for movie id to be randomly created at a time

	data, err := c.store.CreateMovie(movie)
	if err != nil {
		json.NewEncoder(w).Encode(Models.GenericResponse{Code: 404, Status: "FAILURE", Data: "no request body found"})
	}
	response := Models.MovieDetails{Code: 200, Status: "SUCCESS", Data: &data}
	err = json.NewEncoder(w).Encode(response)
	Middlewares.CheckNillError(err)
	return
}

func (c *storeConnector) EditMovieById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Checking if request body is nil or not
	if r.Body == nil {
		err := json.NewEncoder(w).Encode(Models.GenericResponse{Code: 404, Status: "FAILURE", Data: "no request body found"})
		Middlewares.CheckNillError(err)
		return
	}

	//Decoding the request data with struct type movie and checking if it is nil or not
	var movie Models.Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)

	if movie.IsEmpty() {
		err := json.NewEncoder(w).Encode(Models.GenericResponse{Code: 404, Status: "FAILURE", Data: "no data found in json"})
		Middlewares.CheckNillError(err)
		return
	}

	params := mux.Vars(r)

	data, err := c.store.UpdateMovieByID(movie, params)
	response := Models.MovieDetails{Code: 200, Status: "SUCCESS", Data: &data}
	if err != nil {
		err := json.NewEncoder(w).Encode(Models.GenericResponse{Code: 404, Status: "FAILURE", Data: "Invalid"})
		Middlewares.CheckNillError(err)
		return
	}

	err = json.NewEncoder(w).Encode(response)
	Middlewares.CheckNillError(err)
	return
}

func (c *storeConnector) DeleteMovieById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	data, errorS := c.store.DeleteMovieByID(params)
	if errorS != nil {
		err := json.NewEncoder(w).Encode(Models.GenericResponse{Code: 404, Status: "FAILURE", Data: "Id Not Found"})
		Middlewares.CheckNillError(err)
		return
	}

	err := json.NewEncoder(w).Encode(data)
	Middlewares.CheckNillError(err)
	return
}
