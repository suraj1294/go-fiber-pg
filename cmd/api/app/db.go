package app

import (
	"log"

	"github.com/doug-martin/goqu/v9"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/suraj1294/go-fiber-pg-auth/internal/repository"
)

func openDB(dsn string) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func (app *Application) ConnectToDB() error {

	DB, err := openDB(app.Dsn)
	if err != nil {

		return err
	}

	log.Println("Connected to Postgres!")
	appDB := repository.PostgresDB{DB: DB, Goqu: goqu.New("postgres", DB)}

	app.AppDB = &appDB

	return nil

}
