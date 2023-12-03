package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	App "github.com/suraj1294/go-fiber-pg-auth/cmd/api/app"
	"github.com/suraj1294/go-fiber-pg-auth/internal/logger"
)

var defaultPort = "8080"

func main() {

	app := App.GetApplication()

	err := godotenv.Load()
	if err != nil {
		logger.Error("failed to load .env file")
	}

	var (
		dsn        = os.Getenv("DATABASE_URL")
		port       = os.Getenv("PORT")
		jwtSecrete = os.Getenv("JWT_SECRETE")
	)

	//load environment variables

	app.Dsn = dsn
	app.Port = port
	app.JWTIssuer = jwtSecrete

	if app.Port == "" {
		app.Port = defaultPort
	}

	if app.Dsn == "" {
		// read from command line
		flag.StringVar(&app.Dsn, "dsn", "host=localhost port=5432 user=postgres password=postgres dbname=movies sslmode=disable", "Postgres connection string")

	}

	if app.JWTSecrete == "" {
		flag.StringVar(&app.JWTSecrete, "jwt-secrete", "verysecret", "signing secrete")
	}

	flag.StringVar(&app.JWTIssuer, "jwt-issuer", "example.com", "signing issuer")
	flag.StringVar(&app.JWTAudience, "jwt-audience", "example.com", "signing audience")
	//flag.StringVar(&app.CookieDomain, "cookie-domain", "localhost", "cookie domain")
	//flag.StringVar(&app.Domain, "domain", "example.com", "domain")
	flag.Parse()

	//try connection DB and initialize DB repo
	err = app.ConnectToDB()

	if err != nil {
		log.Fatal("failed to connect to DB")
	}

	//initialize Auth
	auth := app.GetAppAuth()
	app.Auth = auth

	//initialize routes
	router := app.RegisterRoutes()
	app.Router = router

	//start server
	app.Router.Listen(fmt.Sprintf(":%s", app.Port))

}
