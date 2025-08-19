package main

import (
	"context"
	"net/http"

	"github.com/EduardoOliveira/notediz/internal/dbkuzu"
	"github.com/EduardoOliveira/notediz/internal/handler"
)

func main() {
	//db := dbsql.MustNew(".")
	//migrations.MustMigrate(context.Background(), db.Db)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	db := dbkuzu.MustNew(ctx, ".")

	handler := handler.New(db)

	http.ListenAndServe(":8080", corsMiddleware(handler.HTTPHandler))
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}
