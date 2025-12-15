package database

import (
	"database/sql"
	"time"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

var db *bun.DB

func GetDbInstance() *bun.DB {
	if db != nil {
		return db
	}

	sqldb := sql.OpenDB(pgdriver.NewConnector(
		pgdriver.WithDSN("postgres://user:password@localhost:5432/dbname?sslmode=disable"), //todo: move connection string to config
	))
	sqldb.SetMaxOpenConns(25)                 // Maximum open connections
	sqldb.SetMaxIdleConns(10)                 // Maximum idle connections
	sqldb.SetConnMaxLifetime(1 * time.Minute) // Connection lifetime
	sqldb.SetConnMaxIdleTime(1 * time.Minute) // Idle connection timeout

	db = bun.NewDB(sqldb, pgdialect.New())

	return db
}
