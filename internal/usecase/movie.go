package usecase

import (
	"context"
	"movie_restfull/internal/entity"
)

func (usecase *movieUsecase) GetAllMovie(ctx context.Context) ([]entity.Movies, error) {
	var movies []entity.Movies

	// Get movie from DB
	movies, err := usecase.movieRepository.GetAllMovie(ctx)
	if err != nil {
		return movies, err
	}

	return movies, nil
}

func (usecase *movieUsecase) GetMovieByTitle(ctx context.Context, title string) ([]entity.Movies, error) {
	var movies []entity.Movies

	// get movie from DB
	movies, err := usecase.movieRepository.GetMovieByTitle(ctx, title)
	if err != nil {
		return movies, err
	}

	return movies, nil
}

func (usecase *movieUsecase) GetMovie(ctx context.Context, id int64) (*entity.Movies, error) {
	// get movie from DB
	movie, err := usecase.movieRepository.GetMovie(ctx, id)
	if err != nil {
		return movie, err
	}

	return movie, nil
}

func (usecase *movieUsecase) CreateMovie(ctx context.Context, movie *entity.Movies) (*entity.Movies, error) {
	var newMovie *entity.Movies

	// Create movie to DB
	id, err := usecase.movieRepository.CreateMovie(ctx, movie)
	if err != nil {
		return newMovie, err
	}

	// Find new movie from DB
	newMovie, err = usecase.movieRepository.GetMovie(ctx, id)
	if err != nil {
		return newMovie, err
	}

	return newMovie, nil
}

func (usecase *movieUsecase) UpdateMovie(ctx context.Context, movie entity.Movies) (entity.Movies, error) {
	var updateMovie *entity.Movies

	// update movie
	err := usecase.movieRepository.UpdateMovie(ctx, movie)
	if err != nil {
		return *updateMovie, err
	}

	// find movie
	updateMovie, err = usecase.movieRepository.GetMovie(ctx, movie.ID)
	if err != nil {
		return *updateMovie, err
	}

	return *updateMovie, nil
}

func (usecase *movieUsecase) DeleteMovie(ctx context.Context, id int64) error {
	// delete from DB
	err := usecase.movieRepository.DeleteMovie(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
