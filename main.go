package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	_ "github.com/pdrum/swagger-automation/docs"

	httpLayer "github.com/varuns-zop/movie/Internal/http/movie"
	serviceLayer "github.com/varuns-zop/movie/Internal/service/movie"
	storeLayer "github.com/varuns-zop/movie/Internal/stores/movie"
)

func getDBObject() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:varun123@/movies")

	if err != nil {
		return nil, errors.New(err.Error())
	} else {
		fmt.Println("DB Connected")
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	return db, nil
}

func main() {

	route := mux.NewRouter()
	db, err := getDBObject()

	if err != nil {
		panic(err)
	}

	storeHandler := storeLayer.New(db)
	serviceable := serviceLayer.New(storeHandler)
	handler := httpLayer.New(serviceable)

	route.HandleFunc("/movies", handler.GetALLMovies).Methods("GET")
	route.HandleFunc("/movie", handler.AddingMovie).Methods("POST")
	route.HandleFunc("/movie/{id}", handler.EditMovieById).Methods("PUT")
	route.HandleFunc("/movie/{id}", handler.DeleteMovieById).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":4000", route))
}
