package db

import (
	"OzonTestTask/OzonTestTask/utils"
	"database/sql"
	_ "embed"
	"log"
	"os"
)

var db *sql.DB = nil

//go:embed init.sql
var initSQL string

// TODO прописать в compose env
func Connect() {
	log.Printf("Connecting to the database")

	username, ok1 := os.LookupEnv("DB_USERNAME")
	password, ok2 := os.LookupEnv("DB_PASSWORD")
	databaseIP, ok3 := os.LookupEnv("DB_HOST")
	databaseName, ok4 := os.LookupEnv("DB_NAME")

	if !ok1 || !ok2 || !ok3 || !ok4 {
		log.Printf("Database environment variables are not set")
		return
	}

	connStr := "user=" + username + " password=" + password + " dbname=" + databaseName + " host=" + databaseIP + " sslmode=disable"

	var err error

	db, err = sql.Open("postgres", connStr)

	utils.ErrorHandler(err, "Error while connecting to the database")

	setupDB()
}

func setupDB() {
	if db == nil {
		log.Printf("Database is not connected")
		return
	}

	_, err := db.Exec(initSQL)

	log.Println("[INFO]: Creating the table")

	utils.ErrorHandler(err, "Error while creating the table")
}
