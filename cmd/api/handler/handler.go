package handler

import (
	"github.com/suraj1294/go-fiber-pg-auth/internal/models"
	"github.com/suraj1294/go-fiber-pg-auth/internal/repository"
)

type AppStore struct {
	UsersStore  *repository.UsersRepo
	MoviesStore *repository.MoviesRepo
	AuthStore   *models.Auth
}

func GetAppHandler(auth models.Auth, appDb repository.PostgresDB) *AppStore {
	moviesRepo := new(repository.MoviesRepo)
	moviesRepo.PostgresDB = appDb

	usersRepo := new(repository.UsersRepo)
	usersRepo.PostgresDB = appDb

	appStore := AppStore{
		MoviesStore: moviesRepo,
		UsersStore:  usersRepo,
		AuthStore:   &auth,
	}

	return &appStore

}
