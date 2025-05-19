package main

import (
	"database/sql"
	"time"

	_ "github.com/lib/pq"
)

const (
	maxOpenDBConn = 25
	maxIdleDBConn = 25
	maxDbLifetime = 5 * time.Minute
)

func initPostgresDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(maxOpenDBConn)
	db.SetConnMaxIdleTime(maxIdleDBConn)
	db.SetConnMaxLifetime(maxDbLifetime)

	return db, nil
}
