package usecase

import (
	"context"
	"movie_restfull/internal/entity"
	repository "movie_restfull/internal/repository"
)

type MovieUsecase interface {
	GetAllMovie(ctx context.Context) ([]entity.Movies, error)
	GetMovie(ctx context.Context, id int64) (*entity.Movies, error)
	CreateMovie(ctx context.Context, movie *entity.Movies) (*entity.Movies, error)
	UpdateMovie(ctx context.Context, movie entity.Movies) (entity.Movies, error)
	DeleteMovie(ctx context.Context, id int64) error
}

type movieUsecase struct {
	movieRepository repository.MovieRepository
}

func NewUsecase(movieRepository repository.MovieRepository) MovieUsecase {
	return &movieUsecase{
		movieRepository: movieRepository,
	}
}
