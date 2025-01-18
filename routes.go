package main

import (
	"github.com/alae-touba/playing-with-go-chi/handlers"
	"github.com/alae-touba/playing-with-go-chi/middlewares"
	"github.com/alae-touba/playing-with-go-chi/services"
	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
	"net/http"
)

func RegisterRoutes(r *chi.Mux, logger *zap.Logger, userService *services.UserService, userHandler *handlers.UserHandler) {
	// public routes
	r.Get("/api/v1/hello_world", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})
	r.Post("/api/v1/users", userHandler.CreateUser)

	// protected routes
	r.Group(func(r chi.Router) {
		r.Use(middlewares.NewAuthMiddleware(logger, userService).BasicAuth)
		r.Get("/api/v1/users/{id}", userHandler.GetUser)
	})
}
