package main

import (
	"OzonTestTask/OzonTestTask/db"
	"OzonTestTask/OzonTestTask/graph"
	"database/sql"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
)

const defaultPort = "8080"

var DB *sql.DB

func main() {
	var err error

	DB, err = db.Connect()
	for err != nil {
		log.Printf("Error connecting to DB: %v", err)
		time.Sleep(5 * time.Second)
		DB, err = db.Connect()
	}

	db.InitDB(DB)

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: graph.NewResolver(DB)}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
