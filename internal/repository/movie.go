package repository

import (
	"context"
	"movie_restfull/internal/entity"
)

func (repo *movieRepository) GetAllMovie(ctx context.Context) ([]entity.Movies, error) {
	return repo.postgres.GetAllMovie(ctx)
}

func (repo *movieRepository) GetMovie(ctx context.Context, id int64) (*entity.Movies, error) {
	return repo.postgres.GetMovie(ctx, id)
}

func (repo *movieRepository) CreateMovie(ctx context.Context, movie *entity.Movies) (int64, error) {
	return repo.postgres.CreateMovie(ctx, movie)
}

func (repo *movieRepository) UpdateMovie(ctx context.Context, movie entity.Movies) error {
	return repo.postgres.UpdateMovie(ctx, movie)
}

func (repo *movieRepository) DeleteMovie(ctx context.Context, id int64) error {
	return repo.postgres.DeleteMovie(ctx, id)
}
