package main

import (
	"github.com/alae-touba/playing-with-go-chi/handlers"
	"github.com/alae-touba/playing-with-go-chi/middlewares"
	"github.com/alae-touba/playing-with-go-chi/services"
	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
	"net/http"
)

func registerV1Routes(r chi.Router, logger *zap.Logger, userService *services.UserService, userHandler *handlers.UserHandler) {
	// public routes
	r.Get("/hello_world", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})
	// Group related endpoints
	r.Route("/users", func(r chi.Router) {
		r.Post("/", userHandler.CreateUser)

		// Protected routes
		r.Group(func(r chi.Router) {
			r.Use(middlewares.NewAuthMiddleware(logger, userService).BasicAuth)
			r.Get("/{id}", userHandler.GetUser)
			r.Patch("/{id}", userHandler.UpdateUser)
			r.Delete("/{id}", userHandler.DeleteUser)

		})
	})
}

func RegisterRoutes(r *chi.Mux, logger *zap.Logger, userService *services.UserService, userHandler *handlers.UserHandler) {
	r.Route("/api/v1", func(r chi.Router) {
		registerV1Routes(r, logger, userService, userHandler)
	})
}
