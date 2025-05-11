package routes

import (
	"net/http"
	"ratelimiter/handlers"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// AuthMiddleware checks for Authorization header
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")

		// Simple validation (Adjust logic for JWT or API key validation)
		if authHeader == "" || authHeader != "Bearer my-secret-token" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func SetupRouter() *chi.Mux {

	r := chi.NewRouter()

	r.Use(middleware.Logger)

	r.Route("/users", func(r chi.Router) {
		r.Post("/", handlers.CreateUser)
		r.Get("/{id}", handlers.GetUserByID)
	})

	return r
}
