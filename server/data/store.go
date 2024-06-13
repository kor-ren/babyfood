package data

import (
	"context"
	"database/sql"
	"net/http"
)

type ctxKey string

const (
	dataContextKey = ctxKey("dataContext")
)

type DataContext struct {
	db *sql.DB
}

func NewDataContext(db *sql.DB) *DataContext {
	dataContext := &DataContext{
		db: db,
	}

	return dataContext
}

func Middleware(db *sql.DB, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		dataContext := NewDataContext(db)

		r = r.WithContext(context.WithValue(r.Context(), dataContextKey, dataContext))

		next.ServeHTTP(w, r)
	})
}

func For(ctx context.Context) *DataContext {
	return ctx.Value(dataContextKey).(*DataContext)
}
