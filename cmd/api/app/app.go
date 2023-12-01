package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/suraj1294/go-fiber-pg-auth/internal/models"
	"github.com/suraj1294/go-fiber-pg-auth/internal/repository"
)

var app Application

type Application struct {
	Router       *fiber.App
	Dsn          string
	Port         string
	AppDB        *repository.PostgresDB
	Auth         *models.Auth
	JWTSecrete   string
	JWTIssuer    string
	JWTAudience  string
	CookieDomain string
	Domain       string
}

func GetApplication() *Application {
	return &app
}
