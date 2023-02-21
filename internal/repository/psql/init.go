package psql

import (
	"context"
	"database/sql"
	"movie_restfull/internal/entity"
)

type MoviePostgres interface {
	GetAllMovie(ctx context.Context) ([]entity.Movies, error)
	GetMovieByTitle(ctx context.Context, title string) ([]entity.Movies, error)
	GetMovie(ctx context.Context, id int64) (*entity.Movies, error)
	CreateMovie(ctx context.Context, movie *entity.Movies) (int64, error)
	UpdateMovie(ctx context.Context, movie entity.Movies) error
	DeleteMovie(ctx context.Context, id int64) error
}

type postgresConnection struct {
	db *sql.DB
}

func NewPostgres(db *sql.DB) MoviePostgres {
	return &postgresConnection{
		db: db,
	}
}
