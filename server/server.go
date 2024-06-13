package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/kor-ren/babyfood/auth"
	"github.com/kor-ren/babyfood/data"
	"github.com/kor-ren/babyfood/graph"
	_ "github.com/mattn/go-sqlite3"
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
	dbPath := getEnvOrDefault("DB_PATH", ":memory:")
	port := getEnvOrDefault("PORT", "8080")
	token := getEnvOrDefault("TOKEN", "test")
	tokenHash := auth.GetTokenHash(token)
	secureCookie := getEnvOrDefault("COOKIE_SECURE", "false")

	staticFilePath := getEnvOrDefault("STATIC_FILES_PATH", "")

	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		panic(fmt.Errorf("could not open db: %w", err))
	}
	defer db.Close()

	initDb(db)

	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

	mux := http.NewServeMux()

	if staticFilePath != "" {
		mux.Handle("/", auth.Middleware(tokenHash, secureCookie, http.FileServer(http.Dir(staticFilePath))))
	}

	if enablePlayground == "true" {
		mux.Handle("/playground", auth.Middleware(tokenHash, secureCookie, playground.Handler("GraphQL playground", "/query")))
	}
	mux.Handle("/query", auth.Middleware(tokenHash, secureCookie, srv))

	mux.HandleFunc("/login", auth.LoginHandler(token, tokenHash, secureCookie))

	middleware := data.Middleware(db, mux)

	log.Printf("starting on http://localhost:%s", port)
	log.Fatal(http.ListenAndServe(":"+port, middleware))
}

func initDb(db *sql.DB) {
	driver, err := sqlite.WithInstance(db, &sqlite.Config{})
	if err != nil {
		log.Panic(fmt.Errorf("could not get driver for db: %w", err))
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://db/migrations",
		"sqlite3",
		driver,
	)

	if err != nil {
		log.Panic(fmt.Errorf("could not get migrations: %w", err))
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Panic(fmt.Errorf("migration failed: %w", err))
	}

	log.Println("Migrations applied successfully!")
}
