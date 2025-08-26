package http

import (
	"crudl/internal/usecase"
	"crudl/pkg/http_server"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/swaggo/http-swagger"
)

func Router(profile *usecase.Profile, c http_server.Config) http.Handler{
	r := chi.NewRouter()

	r.Use(middleware.Recoverer)
	r.Use(middleware.Logger)
	r.Use(func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            w.Header().Set("Access-Control-Allow-Origin", "*")
            w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
            w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
            if r.Method == "OPTIONS" {
                w.WriteHeader(http.StatusOK)
                return
            }
            next.ServeHTTP(w, r)
        })
    })

	r.Get("/swagger/*", httpSwagger.Handler(
    	httpSwagger.URL(c.Swagger),
	))
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "static/index.html")
	})

	handlers := New(profile)

	r.Get("/sub/{id}", handlers.GetOrder)

	return r
}