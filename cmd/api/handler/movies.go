package handler

import (
	"fmt"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/suraj1294/go-fiber-pg-auth/cmd/api/response"
	"github.com/suraj1294/go-fiber-pg-auth/internal/logger"
)

type MoviesHandler struct {
	AppStore
}

func (mh *MoviesHandler) MockMovies(c *fiber.Ctx) error {

	movies := mh.MoviesStore.MockMovies()

	return response.SuccessResponseHandler(c, movies)

}

func (mh *MoviesHandler) Movie(c *fiber.Ctx) error {

	id, err := strconv.Atoi(c.Params("id"))

	if err != nil {
		logger.Error("failed to read id" + err.Error())
		return response.BadRequestHandler(c, "invalid request")

	}

	movie, err := mh.MoviesStore.MovieById(int(id))

	if err != nil {
		return response.NotFoundRequestHandler(c, "movie not found")

	}

	return response.SuccessResponseHandler(c, movie)

}

func (mh *MoviesHandler) AllMovies(c *fiber.Ctx) error {

	fmt.Println(c.Cookies(mh.AuthStore.CookieName))

	movies, err := mh.MoviesStore.Movies()

	if err != nil {
		return response.UnexpectedErrorHandler(c, "failed to get movies")

	}

	return response.SuccessResponseHandler(c, movies)

}
