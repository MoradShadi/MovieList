package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

type Movie struct {
	gorm.Model
	Title       string `json:"title"`
	Description string `json:"description"`
	Watched     bool   `json:"watched"`
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
}

func getDatabaseURL() string {

	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		"postgresdb", "user", "password", "movie_watchlist_db", "5432",
	)
}

func main() {

	db, err := gorm.Open(postgres.Open(getDatabaseURL()), &gorm.Config{})
	if err == nil {
		fmt.Println("Connected to the PostgreSQL database!")
	}

	err = db.AutoMigrate(&Movie{})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to the PostgreSQL database!")

	r := mux.NewRouter()
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Println("Server is running on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))

}

func getMovies(w http.ResponseWriter, r *http.Request) {

	enableCors(&w)
	log.Println("GET /movies endpoint is hit")
	var movies []Movie
	db.Find(&movies)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func createMovie(w http.ResponseWriter, r *http.Request) {

	enableCors(&w)
	log.Println("GET /movies endpoint is hit")
	var movie Movie
	json.NewDecoder(r.Body).Decode(&movie)

	fmt.Println(movie.Title)
	db.Create(&movie)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movie)
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	params := mux.Vars(r)
	var movie Movie
	result := db.First(&movie, params["id"])

	if result.Error != nil {
		http.NotFound(w, r)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movie)
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	params := mux.Vars(r)
	var movie Movie
	result := db.First(&movie, params["id"])

	if result.Error != nil {
		http.NotFound(w, r)
		return
	}

	json.NewDecoder(r.Body).Decode(&movie)
	db.Save(&movie)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movie)
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	params := mux.Vars(r)
	var movie Movie
	result := db.First(&movie, params["id"])

	if result.Error != nil {
		http.NotFound(w, r)
		return
	}

	db.Delete(&movie)

	w.WriteHeader(http.StatusNoContent)
}
