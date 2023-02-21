package handler

import (
	usecase "movie_restfull/internal/usecase"

	"github.com/gin-gonic/gin"
)

type MovieHandler interface {
	GetAllMovie(context *gin.Context)
	GetMovieByTitle(context *gin.Context)
	GetMovie(context *gin.Context)
	CreateMovie(context *gin.Context)
	UpdateMovie(context *gin.Context)
	DeleteMovie(context *gin.Context)
}

type moviehandler struct {
	movieusecase usecase.MovieUsecase
}

func NewHandler(movieusecase usecase.MovieUsecase) MovieHandler {
	return &moviehandler{
		movieusecase: movieusecase,
	}
}
