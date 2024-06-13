package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/dgraph-io/badger/v4"
	"github.com/kor-ren/babyfood/data"
	"github.com/kor-ren/babyfood/graph"
)

func getEnvOrDefault(key string, defaultValue string) string {
	val := os.Getenv(key)

	if val == "" {
		return defaultValue
	}

	return val
}

func main() {

	enablePlayground := getEnvOrDefault("PLAYGROUND_ENABLED", "true")
	dbPath := getEnvOrDefault("DB_PATH", "/tmp/babyfood-db")
	port := getEnvOrDefault("PORT", "8080")

	db, err := badger.Open(badger.DefaultOptions(dbPath))

	if err != nil {
		panic(fmt.Errorf("could not open db: %w", err))
	}

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir("app/dist")))

	if enablePlayground == "true" {
		mux.Handle("/playground", playground.Handler("GraphQL playground", "/query"))
	}
	mux.Handle("/query", srv)

	middleware := data.Middleware(db, mux)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, middleware))
}
