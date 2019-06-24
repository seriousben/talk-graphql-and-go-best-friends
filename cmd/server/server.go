package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/handler"
	"github.com/seriousben/simple-graphql-chat/internal/db"
	"github.com/seriousben/simple-graphql-chat/internal/graph"
	"github.com/seriousben/simple-graphql-chat/internal/resolvers"
)

const defaultPort = "8080"
const defaultDatabaseURL = "postgres://simple-graphql-chat:pass@localhost:5432/simple-graphql-chat?sslmode=disable"

func main() {
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		dbURL = defaultDatabaseURL
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	db, err := db.Dial(dbURL)
	if err != nil {
		log.Fatalf("Error dialing database: %#v", err)
	}
	err = db.AddFixtures()
	if err != nil {
		log.Fatalf("Error add fixtures to database: %#v", err)
	}

	root := resolvers.RootResolver{
		DB: db,
	}

	http.Handle("/", handler.Playground("GraphQL playground", "/query"))
	http.Handle("/query", handler.GraphQL(graph.NewExecutableSchema(graph.Config{Resolvers: &root})))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
