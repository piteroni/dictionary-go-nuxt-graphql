package main

import (
	"net/http"
	"os"

	"piteroni/dictionary-go-nuxt-graphql/graph"
	"piteroni/dictionary-go-nuxt-graphql/graph/generated"
	"piteroni/dictionary-go-nuxt-graphql/pkg/database"
	"piteroni/dictionary-go-nuxt-graphql/pkg/drivers"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

func main() {
	logger := drivers.NewLogger(os.Stdout)

	if err := godotenv.Load(); err != nil {
		logger.Errorf("unexpected error occurred during loading .env: %v", err)
		os.Exit(1)
	}

	db, err := database.ConnectToDatabase()
	if err != nil {
		logger.Errorf("unexpected error occurred during connect database: %v", err)
		os.Exit(1)
	}

	r := &graph.Resolver{
		DB:     db,
		Logger: *logger,
	}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: r}))

	router := mux.NewRouter()

	router.Handle("/api/i/query", srv)
	router.Handle("/", playground.Handler("GraphQL playground", "/api/i/query"))

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:3000"},
		AllowedHeaders: []string{"*"},
		AllowedMethods: []string{
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodPatch,
			http.MethodDelete,
			http.MethodOptions,
			http.MethodHead,
		},
	})

	logger.Error(http.ListenAndServe(":8080", c.Handler(router)))
}
