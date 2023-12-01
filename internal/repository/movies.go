package repository

import (
	"context"
	"time"

	"github.com/doug-martin/goqu/v9"
	_ "github.com/doug-martin/goqu/v9/dialect/postgres"
	_ "github.com/lib/pq"

	"github.com/suraj1294/go-fiber-pg-auth/internal/logger"
	"github.com/suraj1294/go-fiber-pg-auth/internal/models"
)

type MoviesRepo struct {
	PostgresDB
}

func (store *MoviesRepo) MockMovies() *[]models.Movie {
	var movies []models.Movie

	rd, _ := time.Parse("2006-01-02", "1986-03-07")

	highlander := models.Movie{
		ID:          1,
		Title:       "HighLander",
		ReleaseDate: rd,
		MPAARating:  "R",
		RunTime:     116,
		Description: "A nice movie",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	movies = append(movies, highlander)

	return &movies
}

func (store *MoviesRepo) MovieById(id int) (*models.Movie, error) {

	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	movie := store.Goqu.From("movies").Where(goqu.C("id").Eq(id))

	query, _, _ := movie.ToSQL()

	logger.Info(query)

	movies := models.Movie{}
	err := store.DB.GetContext(ctx, &movies, query)

	if err != nil {
		logger.Error("failed to load movie, more info:- " + err.Error())
		return nil, err
	}

	return &movies, nil

}

func (store *MoviesRepo) Movies() (*[]models.Movie, error) {

	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	moviesList := store.Goqu.From("movies")

	query, _, _ := moviesList.ToSQL()

	logger.Info(query)

	movies := []models.Movie{}
	err := store.DB.SelectContext(ctx, &movies, query)

	if err != nil {
		logger.Error("failed to load movies, more info:- " + err.Error())
		return nil, err
	}

	return &movies, nil
}
