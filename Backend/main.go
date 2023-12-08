package main

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gopkg.in/yaml.v2"
)

type DBConfig struct {
	User     string `yaml:"User"`
	Password string `yaml:"Password"`
	Name     string `yaml:"Name"`
	Host     string `yaml:"Host"`
}

type Config struct {
	DB DBConfig `yaml:"DB"`
}

var db *gorm.DB
var config Config

type Movie struct {
	gorm.Model
	Title       string `json:"title"`
	Description string `json:"description"`
	Watched     bool   `json:"watched"`
}

func main() {
	loadConfig()

	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		log.Fatal("DATABASE_URL environment variable not set")
	}

	// Open a connection to the PostgreSQL database using GORM
	db, err := gorm.Open(postgres.Open(databaseURL), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	// Automatically create the 'movies' table based on the 'Movie' struct
	err = db.AutoMigrate(&Movie{})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to the PostgreSQL database!")

	r := gin.Default()

	// Use CORS middleware
	r.Use(cors.Default())

	// Define routes
	r.GET("/movies", getMovies)
	r.POST("/movies", createMovie)
	r.GET("/movies/:id", getMovie)
	r.PUT("/movies/:id", updateMovie)
	r.DELETE("/movies/:id", deleteMovie)

	// Start the server
	fmt.Println("Server is running on :8080")
	log.Fatal(r.Run(":8080"))
}

func loadConfig() {
	file, err := os.Open("config.yaml")
	if err != nil {
		fmt.Println("No configuration file found, using default values.")
		return
	}
	defer file.Close()

	decoder := yaml.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		log.Fatalf("Error decoding configuration: %v", err)
	}
}

func getMovies(c *gin.Context) {
	var movies []Movie
	db.Find(&movies)

	c.JSON(http.StatusOK, movies)
}

func createMovie(c *gin.Context) {
	var movie Movie
	if err := c.ShouldBindJSON(&movie); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	fmt.Println(movie.Title)
	db.Create(&movie)

	c.JSON(http.StatusCreated, movie)
}

func getMovie(c *gin.Context) {
	var movie Movie
	result := db.First(&movie, c.Param("id"))

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Movie not found"})
		return
	}

	c.JSON(http.StatusOK, movie)
}

func updateMovie(c *gin.Context) {
	var movie Movie
	result := db.First(&movie, c.Param("id"))

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Movie not found"})
		return
	}

	if err := c.ShouldBindJSON(&movie); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.Save(&movie)

	c.JSON(http.StatusOK, movie)
}

func deleteMovie(c *gin.Context) {
	var movie Movie
	result := db.First(&movie, c.Param("id"))

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Movie not found"})
		return
	}

	db.Delete(&movie)

	c.Status(http.StatusNoContent)
}
