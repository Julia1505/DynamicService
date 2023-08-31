package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

func connectDB() (*sqlx.DB, error) {
	connString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname = %s sslmode=disable", config.host, config.port, config.user, config.password, config.dbname)
	db, err := sqlx.Open("postgres", connString)
	if err != nil {
		return nil, errors.Wrap(err, "sqlx connect")
	}
	err = db.Ping()
	if err != nil {
		return nil, errors.Wrap(err, "ping no pong")
	}
	return db, nil
}

func closeDB(db *sqlx.DB) {
	db.Close()
}
