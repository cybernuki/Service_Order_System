package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/cybernuki/Service_Order_System/graph"
	"github.com/cybernuki/Service_Order_System/graph/generated"
	"github.com/cybernuki/Service_Order_System/internal/auth"
	"github.com/cybernuki/Service_Order_System/internal/database/models"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
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
	if migration == "" || migration == "True" {
		models.MigrateAll()
	}
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(auth.Middleware())

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{}}))

	router.Handle("/graphql", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	http.ListenAndServe(":"+port, router)

	defer models.CloseDB()
}
