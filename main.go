package main

import (
	"fmt"
	"log"
	"movie_restfull/internal/config"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	// Get the environment
	env := os.Getenv("ENV")
	if env == "production" || env == "staging" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		// Get the config from .env file
		err := godotenv.Load(".env")
		if err != nil {
			log.Fatalf("Error loading .env file")
		}
	}

	// Initialize gin
	r := gin.Default()

	port := os.Getenv("PORT")

	// Load db config
	db, err := config.OpenDB(os.Getenv("POSTGRES_URL"), true)
	if err != nil {
		log.Fatalln(err)
	}
	defer config.CloseDB(db)

	// Init clean arch
	repository := config.InitRepository(db)
	usecase := config.InitUsecase(repository.Allrepository)
	handler := config.InitHandler(usecase.Allusecase)

	// Create the API
	movieRoutes := r.Group("/Movies")
	{
		movieRoutes.GET("/", handler.Allhandler.GetAllMovie)
		movieRoutes.GET("/search/:title", handler.Allhandler.GetMovieByTitle)
		movieRoutes.GET("/:id", handler.Allhandler.GetMovie)
		movieRoutes.POST("/", handler.Allhandler.CreateMovie)
		movieRoutes.PATCH("/:id", handler.Allhandler.UpdateMovie)
		movieRoutes.DELETE(":id", handler.Allhandler.DeleteMovie)
	}

	// Run the gin gonic in port 5000
	runWithPort := fmt.Sprintf("0.0.0.0:%s", port)
	r.Run(runWithPort)
}
