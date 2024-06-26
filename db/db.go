package db

import (
	"database/sql"
	"fmt"
	"log"
	"sync"
)

const (
	DB_USER     = "postgres"
	DB_PASSWORD = "docker"
	DB_NAME     = "gqldemo"
	SSL_MODE    = "disable"
)

var once sync.Once

func Connect() (*sql.DB, error) {
	var db *sql.DB
	var err error
	once.Do(func() {
		dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=%s",
			DB_USER, DB_PASSWORD, DB_NAME, SSL_MODE)
		db, _ = sql.Open("postgres", dbinfo)
		err = db.Ping()
	})
	return db, err
}

func MustExec(db *sql.DB, query string, args ...interface{}) (sql.Result, bool) {
	res, err := db.Exec(query, args...)
	if err != nil {
		log.Printf("Error executing query: %s, %s", query, err)
		return nil, false
	}
	return res, true
}
