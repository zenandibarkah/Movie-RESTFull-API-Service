package config

import (
	repository "movie_restfull/internal/repository"
	usecase "movie_restfull/internal/usecase"
)

type Usecase struct {
	Allusecase usecase.MovieUsecase
}

func InitUsecase(Repository repository.MovieRepository) Usecase {
	return Usecase{
		Allusecase: usecase.NewUsecase(Repository),
	}
}
