package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/handler"
	"github.com/seriousben/simple-graphql-chat/db"
	"github.com/seriousben/simple-graphql-chat/donehookwithsql"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	db, err := db.Dial(dbURL)
	if err != nil {
		log.Fatalf("Error dialing database: %#v", err)
	}

	http.Handle("/", handler.Playground("GraphQL playground", "/query"))
	http.Handle("/query", handler.GraphQL(donehookwithsql.NewExecutableSchema(donehookwithsql.Config{Resolvers: &donehookwithsql.Resolver{db: db}})))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
