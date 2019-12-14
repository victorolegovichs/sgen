package db

import (
	"github.com/jackc/pgx"
)

//Fill in the configuration
func config() (c pgx.ConnConfig) {
	return c
}

func NewConnection() (*pgx.Conn, error) {
	conn, err := pgx.Connect(config())
	if err != nil {
		return nil, err
	}

	return conn, nil
}
