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

type UsersRepo struct {
	PostgresDB
}

func (store *UsersRepo) MockUsers() *[]models.User {
	var users []models.User

	demo_user := models.User{
		ID:        1,
		FirstName: "Suraj",
		LastName:  "Patil",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	users = append(users, demo_user)

	return &users
}

func (store *UsersRepo) GetUserByEmail(email string) (*models.User, error) {

	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	userByEmail := store.Goqu.From("users").Where(goqu.C("email").Eq(email))

	query, _, _ := userByEmail.ToSQL()

	logger.Info(query)

	user := models.User{}
	err := store.DB.GetContext(ctx, &user, query)

	if err != nil {
		logger.Error("failed to get user, more info:- " + err.Error())
		return nil, err
	}

	return &user, nil

}

func (store *UsersRepo) GetUserById(id int) (*models.User, error) {

	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	userByEmail := store.Goqu.From("users").Where(goqu.C("id").Eq(id))

	query, _, _ := userByEmail.ToSQL()

	logger.Info(query)

	user := models.User{}
	err := store.DB.GetContext(ctx, &user, query)

	if err != nil {
		logger.Error("failed to get user, more info:- " + err.Error())
		return nil, err
	}

	return &user, nil

}
