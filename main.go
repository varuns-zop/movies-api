package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/varuns-zop/movie/Controllers"

	"github.com/varuns-zop/movie/Middlewares"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	_ "github.com/pdrum/swagger-automation/docs"
)

func main() {

	db, err := sql.Open("mysql", "root:varun123@/movies")

	if err != nil {
		panic(err)
	} else {
		fmt.Println("DB Connected")
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	Route := mux.NewRouter()

	handler := Middlewares.New(db)
	controllHandler := Controllers.Connect(handler)

	Route.HandleFunc("/movies", controllHandler.GetALLMovies).Methods("GET")
	Route.HandleFunc("/movie", controllHandler.AddingMovie).Methods("POST")
	Route.HandleFunc("/movie/{id}", controllHandler.EditMovieById).Methods("PUT")
	Route.HandleFunc("/movie/{id}", controllHandler.DeleteMovieById).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":4000", Route))
}
