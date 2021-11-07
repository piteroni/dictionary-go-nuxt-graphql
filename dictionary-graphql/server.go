package main

import (
	"encoding/json"
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

const (
	statusOk    = 0
	statusError = 1
)

func main() {
	statusCode := serve()

	os.Exit(statusCode)
}

func serve() int {
	logger := driver.NewLogger(os.Stdout)

	db, err := database.ConnectToDatabase()
	if err != nil {
		logger.Error(err)
		return statusError
	}

	r := &graph.Resolver{
		DB:     db,
		Logger: *logger,
	}

	schema := generated.NewExecutableSchema(generated.Config{Resolvers: r})
	srv := handler.NewDefaultServer(schema)

	router := mux.NewRouter()

	router.HandleFunc("/health", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	router.Handle("/api/i/query", srv)
	router.Handle("/", playground.Handler("GraphQL playground", "/api/i/query"))

	o, err := driver.Env("ALLOW_ORIGINS")
	if err != nil {
		logger.Error(err)
		return statusError
	}

	origins := []string{}
	err = json.Unmarshal([]byte(o), &origins)
	if err != nil {
		logger.Error(err)
		return statusError
	}

	c := cors.New(cors.Options{
		AllowedOrigins: origins,
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
		return statusError
	}

	return statusOk
}
