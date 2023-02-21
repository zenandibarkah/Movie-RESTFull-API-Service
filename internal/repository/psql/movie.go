package psql

import (
	"context"
	"movie_restfull/internal/entity"

	"fmt"

	"time"
)

func (repo *postgresConnection) GetAllMovie(ctx context.Context) ([]entity.Movies, error) {
	query := `
		SELECT id, title, description, rating, image, created_at, update_at 
			FROM movies`

	// Define the contect with 15 timeout
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	var movies []entity.Movies

	rows, err := repo.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var movie entity.Movies

		err := rows.Scan(
			&movie.ID,
			&movie.Title,
			&movie.Description,
			&movie.Rating,
			&movie.Image,
			&movie.CreatedAt,
			&movie.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}

		movies = append(movies, movie)
	}

	return movies, nil
}

func (repo *postgresConnection) GetMovieByTitle(ctx context.Context, title string) ([]entity.Movies, error) {
	query := `
			SELECT id, title, description, rating, image, created_at, update_at 
				FROM movies WHERE lower(title) LIKE CONCAT('%%',$1::text,'%%')`

	// Define the contect with 15 timeout
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	var movies []entity.Movies

	rows, err := repo.db.QueryContext(ctx, query, title)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	// fmt.Println(query)

	for rows.Next() {
		var movie entity.Movies

		err := rows.Scan(
			&movie.ID,
			&movie.Title,
			&movie.Description,
			&movie.Rating,
			&movie.Image,
			&movie.CreatedAt,
			&movie.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}

		movies = append(movies, movie)
	}

	return movies, nil
}

func (repo *postgresConnection) GetMovie(ctx context.Context, id int64) (*entity.Movies, error) {
	query := `
		SELECT id, title, description, rating, image, created_at, update_at 
			FROM movies WHERE id = $1`

	// Define the contect with 15 timeout
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	var movie entity.Movies

	err := repo.db.QueryRowContext(ctx, query, id).Scan(
		&movie.ID,
		&movie.Title,
		&movie.Description,
		&movie.Rating,
		&movie.Image,
		&movie.CreatedAt,
		&movie.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &movie, nil
}

func (repo *postgresConnection) CreateMovie(ctx context.Context, movie *entity.Movies) (int64, error) {
	query := `
		INSERT INTO public.movies (title, description, rating, image, created_at, update_at)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id`

	// Define the contect with 15 timeout
	ctx, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	// Run the query insert
	err := repo.db.QueryRowContext(ctx, query,
		movie.Title,
		movie.Description,
		movie.Rating,
		movie.Image,
		movie.CreatedAt,
		movie.UpdatedAt,
	).Scan(&movie.ID)

	if err != nil {
		return 0, err
	}

	return movie.ID, nil
}

func (repo *postgresConnection) UpdateMovie(ctx context.Context, movie entity.Movies) error {
	query := `UPDATE movies SET title=$1, description=$2, rating=$3, image=$4,
				update_at=$5 WHERE id = $6`

	// Define the contect with 15 timeout
	ctx, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	result, err := repo.db.ExecContext(ctx, query,
		movie.Title,
		movie.Description,
		movie.Rating,
		movie.Image,
		movie.UpdatedAt,
		movie.ID,
	)
	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	fmt.Printf("Affected update : %d", rows)

	return nil
}

func (repo *postgresConnection) DeleteMovie(ctx context.Context, id int64) error {
	query := `DELETE FROM movies WHERE id = $1`

	// Define the contect with 15 timeout
	ctx, cancel := context.WithTimeout(ctx, 15*time.Second)
	defer cancel()

	// Run the delete query
	result, err := repo.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	// Get how many data has been deleted
	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}

	fmt.Printf("Affected delete : %d", rows)
	return nil
}
