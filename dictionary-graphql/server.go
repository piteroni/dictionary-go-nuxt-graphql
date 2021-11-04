package main

import (
	"net/http"
	"os"

	"piteroni/dictionary-go-nuxt-graphql/database"
	"piteroni/dictionary-go-nuxt-graphql/driver"
	"piteroni/dictionary-go-nuxt-graphql/graph"
	"piteroni/dictionary-go-nuxt-graphql/graph/generated"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	logger := driver.NewLogger(os.Stdout)

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

	err = http.ListenAndServe(":8080", c.Handler(router))
	if err != nil {
		logger.Error(err)
	}
}
