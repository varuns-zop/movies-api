package Middlewares

import (
	"fmt"
	"testing"

	"github.com/varuns-zop/movie/Models"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestStoreCreateMovie(t *testing.T) {

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))

	var testcases = []struct {
		expectedOutput Models.Movie
		expectedError  interface{}
		body           Models.Movie
		mockQuerry     []interface{}
	}{
		{
			expectedOutput: Models.Movie{Id: "1", Name: "Dangal", Genre: "BioPic", Rating: 4.5, Plot: "Wrestling for Women", Released: true},
			expectedError:  nil,
			body:           Models.Movie{Name: "Dangal", Genre: "BioPic", Rating: 4.5, Plot: "Wrestling for Women", Released: true},
			mockQuerry: []interface{}{
				mock.ExpectExec("CREATE TABLE IF NOT EXISTS Movies (id INT AUTO_INCREMENT PRIMARY KEY, name VARCHAR(255), genre VARCHAR(255), rating FLOAT, plot VARCHAR(255), released INT);").
					WillReturnResult(sqlmock.NewResult(1, 1)),
				mock.ExpectExec("INSERT INTO Movies(name, genre, rating, plot, released) VALUES (?,?,?,?,?)").
					WithArgs("Dangal", "BioPic", 4.5, "Wrestling for Women", true).
					WillReturnResult(sqlmock.NewResult(1, 1)),
				mock.ExpectQuery("SELECT id,name,released,rating,plot,genre FROM Movies WHERE id = ?").
					WithArgs(1).
					WillReturnRows(sqlmock.NewRows([]string{"id", "name", "released", "rating", "plot", "genre"}).
						AddRow(1, "Dangal", true, 4.5, "Wrestling for Women", "BioPic")),
			},
		},
	}

	CheckError(err, t)

	handler := New(db)

	for i, tt := range testcases {
		result, err := handler.CreateMovie(tt.body)
		if result.Id != tt.expectedOutput.Id || err != nil {
			t.Errorf("Testcase: %v FAILED (Expected Id %v Found Id %v", i+1, tt.expectedOutput.Id, result.Id)
		} else {
			fmt.Println("Testcase:", i+1, " PASSED")
		}
	}

}

func TestStoreGetALlMovies(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	var testcases = []struct {
		expectedOutput []Models.Movie
		expectedError  interface{}
		mockQuerry     []interface{}
	}{
		{
			expectedOutput: []Models.Movie{
				Models.Movie{Id: "1", Name: "Dangal", Genre: "BioPic", Rating: 4.5, Plot: "Wrestling for Women", Released: true},
				Models.Movie{Id: "2", Name: "Ra.One", Genre: "Sci-Fi", Rating: 4.0, Plot: "Science Fiction and Humans", Released: true},
				Models.Movie{Id: "3", Name: "Harry Potter & The Cursed Child", Genre: "Fantasy", Rating: 0.0, Plot: "Witchcraft and Wizardry", Released: false},
			},
			expectedError: nil,
			mockQuerry: []interface{}{
				mock.ExpectQuery("SELECT id,name,released,rating,plot,genre FROM Movies").
					WillReturnRows(sqlmock.NewRows([]string{"id", "name", "released", "rating", "plot", "genre"}).
						AddRow(1, "Dangal", true, 4.5, "Wrestling for Women", "BioPic").
						AddRow(2, "Ra.One", true, 4.0, "Science Fiction and Humans", "Sci-Fi").
						AddRow(3, "Harry Potter & The Cursed Child", false, 0.0, "Witchcraft and Wizardry", "Fantasy")),
			},
		},
	}

	CheckError(err, t)

	handler := New(db)

	for i, tt := range testcases {
		result, err := handler.GetALlMovies()
		if err != nil {
			t.Errorf("testcase %v FAILED - Reason: %s", i, err.Error())
		}
		CheckDeepEqual(result, tt.expectedOutput, i, t)
	}
}

func TestStoreUpdateMovieByID(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	var testcases = []struct {
		expectedOutput interface{}
		expectedError  interface{}
		body           Models.Movie
		mockQuerry     []interface{}
		params         string
	}{
		{
			expectedError:  nil,
			expectedOutput: Models.Movie{Id: "12", Name: "Harry Potter", Genre: "Movie", Rating: 4.3, Plot: "New Plot", Released: true},
			body:           Models.Movie{Genre: "Movie", Rating: 4.3, Plot: "New Plot", Released: true},
			mockQuerry: []interface{}{
				mock.ExpectQuery("UPDATE  Movies SET released=?,rating=?,plot=?,genre=? WHERE id = (?)").
					WithArgs(true, 4.3, "New Plot", "Movie", "12").
					WillReturnRows(sqlmock.NewRows([]string{"id", "name", "genre", "rating", "plot", "released"}).
						AddRow(12, "Harry Potter", "Movie", 4.3, "New Plot", true),
					),
				mock.ExpectQuery("SELECT id,name,released,rating,plot,genre FROM Movies WHERE id = (?)").
					WithArgs("12").
					WillReturnRows(sqlmock.NewRows([]string{"id", "name", "genre", "rating", "plot", "released"}).
						AddRow(12, "Harry Potter", true, 4.3, "New Plot", "Movie"),
					),
			},
			params: "12",
		},
	}

	CheckError(err, t)

	handler := New(db)

	for i, tt := range testcases {
		param := make(map[string]string)
		param["id"] = tt.params
		result, err := handler.UpdateMovieByID(tt.body, param)
		CheckError(err, t)
		CheckDeepEqual(result, tt.expectedOutput, i, t)
	}
}

func TestStoreDeleteMovieByID(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	var testcases = []struct {
		expectedOutput interface{}
		expectedError  interface{}
		body           Models.Movie
		mockQuerry     []interface{}
		params         string
	}{
		{
			expectedError:  nil,
			expectedOutput: Models.GenericResponse{Code: 200, Status: "SUCCESS", Data: "Movie deleted successfully."},
			mockQuerry: []interface{}{
				mock.ExpectQuery("DELETE FROM Movies WHERE id = (?)").
					WithArgs("12").WillReturnRows(sqlmock.NewRows([]string{})),
			},
			params: "12",
		},
		{
			expectedError:  Models.GenericResponse{Code: 404, Status: "FAILURE", Data: "ID Not Found."},
			expectedOutput: nil,
			mockQuerry: []interface{}{
				mock.ExpectQuery("DELETE FROM Movies WHERE id = (?)").
					WithArgs("12").WillReturnRows(sqlmock.NewRows([]string{})),
			},
			params: "13",
		},
	}

	if err != nil {
		t.Errorf("SQL mock error %s", err.Error())
	}

	handler := New(db)

	for i, tt := range testcases {
		param := make(map[string]string)
		param["id"] = tt.params
		result, err := handler.DeleteMovieByID(param)

		if tt.expectedError != nil {
			CheckDeepEqual(result, tt.expectedError, i, t)
		} else {
			CheckError(err, t)
		}

		if tt.expectedOutput != nil {
			CheckDeepEqual(result, tt.expectedOutput, i, t)
		}
	}
}
