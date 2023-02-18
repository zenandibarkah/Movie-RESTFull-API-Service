package config

import (
	handler "movie_restfull/internal/handler"
	usecase "movie_restfull/internal/usecase"
)

type Handler struct {
	Allhandler handler.MovieHandler
}

func InitHandler(Usecase usecase.MovieUsecase) Handler {
	return Handler{
		Allhandler: handler.NewHandler(Usecase),
	}
}
