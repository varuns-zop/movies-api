package Controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"

	"github.com/varuns-zop/movie/Models"
)

func TestAddingMovie(t *testing.T) {
	var testcase = []struct {
		expectedOutput Models.MovieDetails
		expectedError  error
		body           Models.Movie
	}{
		{
			expectedError: nil, body: Models.Movie{Name: "Independence Day", Genre: "Sci-Fi", Rating: 4.5, Plot: "", Released: true},
			expectedOutput: Models.MovieDetails{Code: 200, Status: "SUCCESS", Data: &Models.Movie{Id: "10", Name: "Independence Day", Genre: "Sci-Fi", Rating: 4.5, Plot: "", Released: true}},
		},
		{
			expectedError: nil, body: Models.Movie{Name: "Avatar", Genre: "Sci-Fi", Rating: 4.9, Plot: "", Released: true},
			expectedOutput: Models.MovieDetails{Code: 200, Status: "SUCCESS", Data: &Models.Movie{Id: "10", Name: "Avatar", Genre: "Sci-Fi", Rating: 4.9, Plot: "", Released: true}},
		},
		{
			expectedError: nil, body: Models.Movie{Name: "Harry Potter", Genre: "Fantasy", Rating: 4.9, Plot: "", Released: true},
			expectedOutput: Models.MovieDetails{Code: 200, Status: "SUCCESS", Data: &Models.Movie{Id: "10", Name: "Harry Potter", Genre: "Fantasy", Rating: 4.9, Plot: "", Released: true}},
		},
	}

	for i, tt := range testcase {

		data, _ := json.Marshal(tt.body)
		w := httptest.NewRecorder()
		r := httptest.NewRequest("POST", "/movie", bytes.NewReader(data))
		AddingMovie(w, r)

		res := w.Result()

		data, err := io.ReadAll(res.Body)
		checkNilError(err, t)

		parsed, err := json.Marshal(tt.expectedOutput)
		checkNilError(err, t)

		if string(data[:len(data)-1]) != string(parsed) {
			t.Errorf("expected %v got %v", string(parsed), string(data[:len(data)-1]))
		} else {
			fmt.Println("POST METHOD -> Testcase ", i+1, ": PASSED ✓")
		}
	}
}

func TestEditMovieById(t *testing.T) {

	var testcase = []struct {
		expectedOutput Models.MovieDetails
		expectedError  error
		body           Models.Movie
	}{
		{
			expectedError: nil, body: Models.Movie{Name: "Independence Day", Genre: "Sci-Fi", Rating: 4.5, Plot: "", Released: true},
			expectedOutput: Models.MovieDetails{Code: 200, Status: "SUCCESS", Data: &Models.Movie{Id: "10", Name: "Independence Day", Genre: "Sci-Fi", Rating: 4.5, Plot: "", Released: true}},
		},
		{
			expectedError: nil, body: Models.Movie{Name: "Avatar", Genre: "Sci-Fi", Rating: 4.9, Plot: "", Released: true},
			expectedOutput: Models.MovieDetails{Code: 200, Status: "SUCCESS", Data: &Models.Movie{Id: "10", Name: "Avatar", Genre: "Sci-Fi", Rating: 4.9, Plot: "", Released: true}},
		},
		{
			expectedError: nil, body: Models.Movie{Name: "Harry Potter", Genre: "Fantasy", Rating: 4.9, Plot: "", Released: true},
			expectedOutput: Models.MovieDetails{Code: 200, Status: "SUCCESS", Data: &Models.Movie{Id: "10", Name: "Harry Potter", Genre: "Fantasy", Rating: 4.9, Plot: "", Released: true}},
		},
	}

	for i, tt := range testcase {
		data, _ := json.Marshal(tt.body)
		w := httptest.NewRecorder()
		r := httptest.NewRequest("PUT", "/movie", bytes.NewReader(data))

		params := map[string]string{
			"id": "10",
		}

		r = mux.SetURLVars(r, params)
		EditMovieById(w, r)

		res := w.Result()

		data, err := io.ReadAll(res.Body)
		checkNilError(err, t)

		parsed, err := json.Marshal(tt.expectedOutput)
		checkNilError(err, t)

		fmt.Println(string(data), string(parsed))

		if string(data[:len(data)-1]) != string(parsed) {
			t.Errorf("expected %v got %v", string(parsed), string(data))
		} else {
			fmt.Println("PUT METHOD -> Testcase ", i+1, ": PASSED ✓")
		}

	}
}

func TestGetAllMovies(t *testing.T) {
	var testcase = []struct {
		expectedOutput []Models.Movie
		expectedError  error
	}{
		{
			expectedError: nil,
			expectedOutput: []Models.Movie{
				Models.Movie{Id: "10", Name: "Independence Day", Genre: "Sci-Fi", Rating: 4.5, Plot: "", Released: true},
				Models.Movie{Id: "10", Name: "Avatar", Genre: "Sci-Fi", Rating: 4.9, Plot: "", Released: true},
				Models.Movie{Id: "10", Name: "Harry Potter", Genre: "Fantasy", Rating: 4.9, Plot: "", Released: true},
			},
		},
	}

	for i, tt := range testcase {

		w := httptest.NewRecorder()
		r := httptest.NewRequest("GET", "/movies", nil)
		GetAllMovies(w, r)

		res := w.Result()

		data, err := io.ReadAll(res.Body)
		checkNilError(err, t)

		parsed, err := json.Marshal(tt.expectedOutput)
		checkNilError(err, t)

		if string(data[:len(data)-1]) != string(parsed) {
			t.Errorf("expected %v got %v", string(parsed), string(data[:len(data)-1]))
		} else {
			fmt.Println("GET METHOD -> Testcase ", i+1, ": PASSED ✓")
		}
	}
}

func TestDeleteMovieById(t *testing.T) {
	var testcase = []struct {
		expectedOutput Models.GenericResponse
		expectedError  error
	}{
		{
			expectedError:  nil,
			expectedOutput: Models.GenericResponse{Code: 200, Status: "SUCCESS", Data: "Movie deleted successfully."},
		},
	}

	for i, tt := range testcase {

		w := httptest.NewRecorder()
		r := httptest.NewRequest("DELETE", "/movie", nil)

		params := map[string]string{
			"id": "10",
		}

		r = mux.SetURLVars(r, params)
		DeleteMovieById(w, r)

		res := w.Result()

		data, err := io.ReadAll(res.Body)
		checkNilError(err, t)

		parsed, err := json.Marshal(tt.expectedOutput)
		checkNilError(err, t)

		if string(data[:len(data)-1]) != string(parsed) {
			t.Errorf("expected %v got %v", string(parsed), string(data[:len(data)-1]))
		} else {
			fmt.Println("DELETE METHOD -> Testcase ", i+1, ": PASSED ✓")
		}

	}
}

func TestGetSingleMovieById(t *testing.T) {

	var testcase = []struct {
		expectedOutput Models.MovieDetails
		expectedError  error
		body           Models.Movie
	}{
		{
			expectedError: nil, body: Models.Movie{Name: "Avatar", Genre: "Sci-Fi", Rating: 4.9, Plot: "", Released: true},
			expectedOutput: Models.MovieDetails{Code: 200, Status: "SUCCESS", Data: &Models.Movie{Id: "10", Name: "Avatar", Genre: "Sci-Fi", Rating: 4.9, Plot: "", Released: true}},
		},
	}

	for i, tt := range testcase {
		data, _ := json.Marshal(tt.body)
		w := httptest.NewRecorder()
		r := httptest.NewRequest("GET", "/movie", bytes.NewReader(data))

		params := map[string]string{
			"id": "10",
		}

		r = mux.SetURLVars(r, params)
		GetSingleMovieById(w, r)

		res := w.Result()

		data, err := io.ReadAll(res.Body)
		checkNilError(err, t)

		parsed, err := json.Marshal(tt.expectedOutput)
		checkNilError(err, t)

		if string(data[:len(data)-1]) != string(parsed) {
			t.Errorf("expected %v got %v", string(parsed), string(data[:len(data)-1]))
		} else {
			fmt.Println("GET BY ID METHOD  -> Testcase ", i+1, ": PASSED ✓")
		}

	}
}

func checkNilError(err error, t *testing.T) {
	if err != nil {
		t.Errorf("expected error to be nil got %v", err)
	}
}
