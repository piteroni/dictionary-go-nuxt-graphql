package main

import (
	"context"
	"encoding/json"
	"net/http"
	"os"
	"strings"

	"piteroni/dictionary-go-nuxt-graphql/database"
	"piteroni/dictionary-go-nuxt-graphql/driver"
	"piteroni/dictionary-go-nuxt-graphql/graph"
	"piteroni/dictionary-go-nuxt-graphql/graph/generated"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
)

var logger = driver.NewLogger(os.Stdout)

func main() {
	logger.Error(serve())
}

func serve() error {
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

	srv.SetRecoverFunc(func(ctx context.Context, err interface{}) error {
		logger.Printf("unexpected error: %#v", err)

		return errors.New("Internal server error!!")
	})

	router := mux.NewRouter()

	router.HandleFunc("/health", func(rw http.ResponseWriter, _ *http.Request) {
		rw.WriteHeader(http.StatusOK)
	})

	router.Handle("/graphql", srv)
	router.Handle("/", playground.Handler("GraphQL playground", "/graphql"))

	c, err := cors()
	if err != nil {
		return err
	}

	router.Use(c)

	err = http.ListenAndServe(":8080", router)
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func cors() (func(http.Handler) http.Handler, error) {
	o, err := driver.Env("ALLOW_ORIGINS")
	if err != nil {
		return nil, err
	}

	origins := []string{}
	err = json.Unmarshal([]byte(o), &origins)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	methods := []string{
		http.MethodGet,
		http.MethodPost,
		http.MethodPut,
		http.MethodOptions,
		http.MethodPatch,
		http.MethodTrace,
	}

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", strings.Join(origins, ","))
			w.Header().Set("Access-Control-Allow-Headers", "*")
			w.Header().Set("Access-Control-Allow-Methods", strings.Join(methods, ","))

			next.ServeHTTP(w, r)
		})
	}, nil
}
