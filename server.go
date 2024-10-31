package main

import (
	"context"
	"go-cms-gql/database"
	"go-cms-gql/directives"
	"go-cms-gql/graph"
	"go-cms-gql/graph/middlewares"
	"go-cms-gql/utils"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

const defaultPort = "8080"

type operation func(ctx context.Context) error

func NewGraphQLHandler() *chi.Mux {
	var router *chi.Mux = chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middlewares.NewMiddleware())

	directives.InitValidator()

	c := graph.Config{
		Resolvers: graph.InitResolver(),
	}

	c.Directives.Validate = directives.ValidateRequest
	c.Directives.Admin = directives.CheckAdmin
	c.Directives.Auth = directives.GetAuthenticatedUser

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

	server := &http.Server{
		Addr:    ":" + port,
		Handler: handler,
	}

	go func() {
		log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server error: %v", err)
		}
	}()

	// Setup graceful shutdown
	ctx := context.Background()
	timeout := 2 * time.Second

	wait := gracefulShutdown(ctx, timeout, map[string]operation{
		"database": func(ctx context.Context) error {
			return database.Disconnect(ctx)
		},
		"http-server": func(ctx context.Context) error {
			shutdownCtx, cancel := context.WithTimeout(ctx, timeout)
			defer cancel()
			return server.Shutdown(shutdownCtx)
		},
	})

	<-wait
	log.Println("All resources have been shut down gracefully.")
}

// gracefulShutdown performs application shut down gracefully.
func gracefulShutdown(ctx context.Context, timeout time.Duration, ops map[string]operation) <-chan struct{} {
	wait := make(chan struct{})
	go func() {
		s := make(chan os.Signal, 1)

		// add any other syscalls that you want to be notified with
		signal.Notify(s, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
		<-s

		log.Println("Initiating shutdown...")

		// set timeout for the ops to be done to prevent system hang
		timeoutFunc := time.AfterFunc(timeout, func() {
			log.Printf("Timeout of %d ms has elapsed, forcing exit", timeout.Milliseconds())
			os.Exit(0)
		})

		defer timeoutFunc.Stop()

		var wg sync.WaitGroup

		// Execute shutdown operations asynchronously to save time
		for key, op := range ops {
			wg.Add(1)
			innerOp := op
			innerKey := key
			go func() {
				defer wg.Done()

				log.Printf("Shutting down: %s", innerKey)
				if err := innerOp(ctx); err != nil {
					log.Printf("%s: shutdown failed: %s", innerKey, err.Error())
					return
				}

				log.Printf("%s has been shut down gracefully", innerKey)
			}()
		}

		wg.Wait()

		close(wait)
	}()

	return wait
}
