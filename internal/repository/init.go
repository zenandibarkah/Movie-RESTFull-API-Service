package repository

import (
	"context"
	"database/sql"
	"movie_restfull/internal/entity"
	"movie_restfull/internal/repository/psql"
)

type MovieRepository interface {
	GetAllMovie(ctx context.Context) ([]entity.Movies, error)
	GetMovie(ctx context.Context, id int64) (*entity.Movies, error)
	CreateMovie(ctx context.Context, movie *entity.Movies) (int64, error)
	UpdateMovie(ctx context.Context, movie entity.Movies) error
	DeleteMovie(ctx context.Context, id int64) error
}

type movieRepository struct {
	postgres psql.MoviePostgres
}

func NewRepository(db *sql.DB) MovieRepository {
	return &movieRepository{
		postgres: psql.NewPostgres(db),
	}
}
