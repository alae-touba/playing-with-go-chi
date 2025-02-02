package main

import (
	"github.com/alae-touba/playing-with-go-chi/handlers"
	"github.com/alae-touba/playing-with-go-chi/middlewares"
	"github.com/alae-touba/playing-with-go-chi/services"
	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
	"net/http"
)

func registerV1Routes(r chi.Router, logger *zap.Logger, userService *services.UserService, 
	userHandler *handlers.UserHandler, topicHandler *handlers.TopicHandler) {
	// public routes
	r.Get("/hello_world", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})

	// users routes
	r.Route("/users", func(r chi.Router) {
		// public
		r.Post("/", userHandler.CreateUser)

		// protected
		r.Group(func(r chi.Router) {
			r.Use(middlewares.NewAuthMiddleware(logger, userService).BasicAuth)
			r.Get("/", userHandler.GetUsers)
			r.Get("/{id}", userHandler.GetUser)
			r.Patch("/{id}", userHandler.UpdateUser)
			r.Delete("/{id}", userHandler.DeleteUser)
		})
	})

	// topics routes
	r.Route("/topics", func(r chi.Router) {	
		// protected
		r.Group(func(r chi.Router) {
			r.Use(middlewares.NewAuthMiddleware(logger, userService).BasicAuth)
			r.Get("/", topicHandler.GetTopics)
			r.Get("/{id}", topicHandler.GetTopic)
			r.Post("/", topicHandler.CreateTopic)
			r.Patch("/{id}", topicHandler.UpdateTopic)
			r.Delete("/{id}", topicHandler.DeleteTopic)
		})
	})
}

func RegisterRoutes(r *chi.Mux, logger *zap.Logger, userService *services.UserService, userHandler *handlers.UserHandler, 
	topicHandler *handlers.TopicHandler) {
	r.Route("/api/v1", func(r chi.Router) {
		registerV1Routes(r, logger, userService, userHandler, topicHandler)
	})
}
