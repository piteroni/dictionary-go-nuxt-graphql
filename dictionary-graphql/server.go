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
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
)

var logger *driver.AppLogger

func init() {
	logger = driver.NewLogger(os.Stdout)
}

func main() {
	logger.Error(serve())
}

func serve() error {
	o, err := driver.Env("ALLOW_ORIGINS")
	if err != nil {
		return err
	}

	origins := []string{}
	err = json.Unmarshal([]byte(o), &origins)
	if err != nil {
		return errors.WithStack(err)
	}

	allowHeaders := handlers.AllowedHeaders([]string{"*"})
	allowOrigins := handlers.AllowedOrigins(origins)
	allowMethods := handlers.AllowedMethods([]string{
		http.MethodGet,
		http.MethodPost,
		http.MethodPut,
		http.MethodPatch,
		http.MethodDelete,
		http.MethodOptions,
		http.MethodHead,
	})

	db, err := database.ConnectToDatabase()
	if err != nil {
		return err
	}

	r := &graph.Resolver{
		DB:     db,
		Logger: logger,
	}

	schema := generated.NewExecutableSchema(generated.Config{Resolvers: r})
	srv := handler.NewDefaultServer(schema)

	router := mux.NewRouter()

	router.HandleFunc("/health", func(rw http.ResponseWriter, _ *http.Request) {
		rw.WriteHeader(http.StatusOK)
	})

	router.Handle("/api/i/query", srv)
	router.Handle("/", playground.Handler("GraphQL playground", "/api/i/query"))

	err = http.ListenAndServe(":8080", handlers.CORS(allowHeaders, allowOrigins, allowMethods)(router))
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}
