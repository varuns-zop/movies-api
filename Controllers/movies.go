package Controllers

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/varuns-zop/movie/Models"
)

var movies []Models.Movie

func PopulateMovieData() {
	movies = append(movies, Models.Movie{Id: "908", Name: "Silicon valley", Genre: "Comedy", Rating: "4.5", Plot: "Richard, a programmer, creates an app called the Pied Piper and tries to get\ninvestors for it. Meanwhile, five other programmers struggle to make their mark in Silicon\nValley.", Released: true})
	movies = append(movies, Models.Movie{Id: "323", Name: "Independence Day", Genre: "Sci-Fi", Rating: "4.9", Plot: "Richard, a programmer, creates an app called the Pied Piper and tries to get\ninvestors for it. Meanwhile, five other programmers struggle to make their mark in Silicon\nValley.", Released: true})
	movies = append(movies, Models.Movie{Id: "543", Name: "Imitation Game", Genre: "Thriller", Rating: "4.1", Plot: "Richard, a programmer, creates an app called the Pied Piper and tries to get\ninvestors for it. Meanwhile, five other programmers struggle to make their mark in Silicon\nValley.", Released: true})
	movies = append(movies, Models.Movie{Id: "112", Name: "Into the Woods", Genre: "Horror", Rating: "4.9", Plot: "Richard, a programmer, creates an app called the Pied Piper and tries to get\ninvestors for it. Meanwhile, five other programmers struggle to make their mark in Silicon\nValley.", Released: true})
}

func GetAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func GetSingleMovieById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// getting the query parameter from the url endpoint
	params := mux.Vars(r)

	for _, movie := range movies {
		if movie.Id == params["id"] {
			res := Models.MovieDetails{Code: 200, Status: "SUCCESS", Data: &movie}
			json.NewEncoder(w).Encode(res)
			return
		}
	}
	json.NewEncoder(w).Encode(Models.GenericResponse{Code: 404, Status: "FAILURE", Data: "no movie found with id"})
}

func AddingMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Checking if request body is nil or not
	if r.Body == nil {
		json.NewEncoder(w).Encode(Models.GenericResponse{Code: 404, Status: "FAILURE", Data: "no request body found"})
	}

	//Decoding the request data with struct type movie and checking if it is nil or not
	var movie Models.Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	if movie.IsEmpty() {
		json.NewEncoder(w).Encode(Models.GenericResponse{Code: 404, Status: "FAILURE", Data: "no data found in json"})
		return
	}

	// creating the random value for movie Id to be randomly created at a time
	rand.Seed(time.Now().UnixNano())
	movie.Id = strconv.Itoa(rand.Intn(1000))
	movies = append(movies, movie)

	// sending back the response with newly created movie
	res := Models.MovieDetails{Code: 200, Status: "SUCCESS", Data: &movie}
	json.NewEncoder(w).Encode(res)
	return
}

func EditMovieById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for index, movie := range movies {
		if movie.Id == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			var movie Models.Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.Id = params["id"]
			movies = append(movies, movie)
			res := Models.MovieDetails{Code: 200, Status: "SUCCESS", Data: &movie}
			json.NewEncoder(w).Encode(res)
			return
		}
	}
	json.NewEncoder(w).Encode(Models.GenericResponse{Code: 404, Status: "FAILURE", Data: "Id Not found"})
	return

}

func DeleteMovieById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for index, movie := range movies {
		if movie.Id == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			json.NewEncoder(w).Encode(Models.GenericResponse{Code: 200, Status: "SUCCESS", Data: "Movie deleted successfully."})
			break
		}
	}
	json.NewEncoder(w).Encode(Models.GenericResponse{Code: 404, Status: "FAILURE", Data: "no movie found with id"})
}
