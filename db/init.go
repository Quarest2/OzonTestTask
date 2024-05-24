package db

import (
	"database/sql"
)

func InitDB(db *sql.DB) {
	MustExec(db, "DROP TABLE IF EXISTS posts")
	MustExec(db, "DROP TABLE IF EXISTS comments")
	MustExec(db, "CREATE TABLE public.posts (id SERIAL PRIMARY KEY, author TEXT, commentsAllowed BOOLEAN)")
	MustExec(db, "CREATE TABLE public.comments (id SERIAL PRIMARY KEY, author TEXT, parentPost SERIAL, parentComment SERIAL)")
}
