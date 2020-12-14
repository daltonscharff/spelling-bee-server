package db

import (
	"fmt"

	"github.com/daltonscharff/spelling-bee-server/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func getConnectionString(c config.Config) string {
	return fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		c.Database.Host, c.Database.Port, c.Database.Username, c.Database.Password, c.Database.Name)
}

func Connect(config config.Config) (*sqlx.DB, error) {
	db, err := sqlx.Connect("postgres", getConnectionString(config))
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
