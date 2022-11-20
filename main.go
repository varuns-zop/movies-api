package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/pdrum/swagger-automation/docs"
	"github.com/varuns-zop/movie/Controllers"
)

func main() {

	Route := mux.NewRouter()

	Route.HandleFunc("/movies", Controllers.GetAllMovies).Methods("GET")
	Route.HandleFunc("/movie/{id}", Controllers.GetSingleMovieById).Methods("GET")
	Route.HandleFunc("/movie", Controllers.AddingMovie).Methods("POST")
	Route.HandleFunc("/movie/{id}", Controllers.EditMovieById).Methods("PUT")
	Route.HandleFunc("/movie/{id}", Controllers.DeleteMovieById).Methods("DELETE")

	Controllers.PopulateMovieData()

	log.Fatal(http.ListenAndServe(":4000", Route))
}
