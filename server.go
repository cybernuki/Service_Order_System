package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/cybernuki/Service_Order_System/graph"
	"github.com/cybernuki/Service_Order_System/graph/generated"
	"github.com/cybernuki/Service_Order_System/internal/database/models"
)

const defaultPort = "8000"

func main() {
	// Getting user listening port
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	if models.InitDB(models.DBConfigParams{}) != nil {
		os.Exit(-127)
	}

	// Getting user migration requeriment
	migration := os.Getenv("MIGRATE")
	if migration == "true" || migration == "True" {
		models.MigrateAll()
	}

	if true {
		srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

		http.Handle("/", playground.Handler("GraphQL playground", "/query"))
		http.Handle("/query", srv)

		log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
		http.ListenAndServe(":"+port, nil)
	}

	defer models.CloseDB()
}
