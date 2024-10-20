package main

import (
	"go-cms-gql/database"
	"go-cms-gql/graph"
	"go-cms-gql/graph/middlewares"
	"go-cms-gql/utils"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

const defaultPort = "8080"

func NewGraphQLHandler() *chi.Mux {
	var router *chi.Mux = chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middlewares.NewMiddleware())

	utils.InitValidator()

	c := graph.Config{
		Resolvers: graph.InitResolver(),
	}

	c.Directives.Validate = utils.ValidateRequest

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(c))

	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)

	return router
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	err := database.Connect(utils.GetValue("DATABASE_NAME"))
	if err != nil {
		log.Fatalf("Cannot connect to the database: %v\n", err)
	}

	var handler *chi.Mux = NewGraphQLHandler()

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, handler))
}
