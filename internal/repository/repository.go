package repository

import (
	"time"

	"github.com/doug-martin/goqu/v9"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const dbTimeout = time.Second * 3

type PostgresDB struct {
	DB   *sqlx.DB
	Goqu *goqu.Database
}
