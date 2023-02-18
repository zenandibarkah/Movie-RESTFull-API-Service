package config

import (
	"database/sql"

	repository "movie_restfull/internal/repository"
)

type Repository struct {
	Allrepository repository.MovieRepository
}

func InitRepository(db *sql.DB) Repository {
	return Repository{
		Allrepository: repository.NewRepository(db),
	}
}
