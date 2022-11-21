package Controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/varuns-zop/movie/Middlewares"
	"github.com/varuns-zop/movie/Models"
)

var movies []Models.Movie

func PopulateMovieData() {
	movies = append(movies, Models.Movie{Id: "10", Name: "Silicon valley", Genre: "Comedy", Rating: 4.5, Plot: "Richard, a programmer, creates an app called the Pied Piper and tries to get\ninvestors for it. Meanwhile, five other programmers struggle to make their mark in Silicon\nValley.", Released: true})
	movies = append(movies, Models.Movie{Id: "323", Name: "Independence Day", Genre: "Sci-Fi", Rating: 4.9, Plot: "Richard, a programmer, creates an app called the Pied Piper and tries to get\ninvestors for it. Meanwhile, five other programmers struggle to make their mark in Silicon\nValley.", Released: true})
	movies = append(movies, Models.Movie{Id: "543", Name: "Imitation Game", Genre: "Thriller", Rating: 4.1, Plot: "Richard, a programmer, creates an app called the Pied Piper and tries to get\ninvestors for it. Meanwhile, five other programmers struggle to make their mark in Silicon\nValley.", Released: true})
	movies = append(movies, Models.Movie{Id: "112", Name: "Into the Woods", Genre: "Horror", Rating: 4.9, Plot: "Richard, a programmer, creates an app called the Pied Piper and tries to get\ninvestors for it. Meanwhile, five other programmers struggle to make their mark in Silicon\nValley.", Released: true})
}

func GetAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(movies)
	Middlewares.CheckNillError(err)
}

func GetSingleMovieById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// getting the query parameter from the url endpoint
	params := mux.Vars(r)

	for _, movie := range movies {
		if movie.Id == params["id"] {
			res := Models.MovieDetails{Code: 200, Status: "SUCCESS", Data: &movie}
			err := json.NewEncoder(w).Encode(res)
			Middlewares.CheckNillError(err)
			return
		}
	}
	err := json.NewEncoder(w).Encode(Models.GenericResponse{Code: 404, Status: "FAILURE", Data: "no movie found with id"})
	Middlewares.CheckNillError(err)
}

func AddingMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Checking if request body is nil or not
	if r.Body == nil {
		err := json.NewEncoder(w).Encode(Models.GenericResponse{Code: 404, Status: "FAILURE", Data: "no request body found"})
		Middlewares.CheckNillError(err)
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
	//rand.Seed(time.Now().UnixNano())
	//movie.Id = strconv.Itoa(rand.Intn(1000))
	movie.Id = "10"
	movies = append(movies, movie)

	// sending back the response with newly created movie
	res := Models.MovieDetails{Code: 200, Status: "SUCCESS", Data: &movie}
	err := json.NewEncoder(w).Encode(res)
	Middlewares.CheckNillError(err)
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
			err := json.NewEncoder(w).Encode(res)
			Middlewares.CheckNillError(err)
			return
		}
	}
	err := json.NewEncoder(w).Encode(Models.GenericResponse{Code: 404, Status: "FAILURE", Data: "Id Not found"})
	Middlewares.CheckNillError(err)
	return

}

func DeleteMovieById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	// Iterating through the movies to delete the movie of particular index
	for index, movie := range movies {
		if movie.Id == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			err := json.NewEncoder(w).Encode(Models.GenericResponse{Code: 200, Status: "SUCCESS", Data: "Movie deleted successfully."})
			Middlewares.CheckNillError(err)
			break
		}
	}
}
