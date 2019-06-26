package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/handler"
	"github.com/seriousben/talk-graphql/db"
	"github.com/seriousben/talk-graphql/donegenallresolvers"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	db, err := db.DialWithEnv()
	if err != nil {
		log.Fatalf("Error dialing database: %#v", err)
	}

	err = db.AddFixtures()
	if err != nil {
		log.Fatalf("Error adding fixtures: %#v", err)
	}

	http.Handle("/", handler.Playground("GraphQL playground", "/query"))
	http.Handle("/query", handler.GraphQL(donegenallresolvers.NewExecutableSchema(donegenallresolvers.Config{Resolvers: &donegenallresolvers.Resolver{DB: db}})))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
