package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/alae-touba/playing-with-go-chi/constants"
	"github.com/alae-touba/playing-with-go-chi/handlers"
	"github.com/alae-touba/playing-with-go-chi/middlewares"
	"github.com/alae-touba/playing-with-go-chi/services"
	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
)

func configureLogger() *zap.Logger {
	// Read the zap configuration from the external JSON file
	configFile, err := os.ReadFile("zap_config.json")
	if err != nil {
		panic(fmt.Sprintf("Failed to read zap config file: %v", err))
	}

	var cfg zap.Config
	if err := json.Unmarshal(configFile, &cfg); err != nil {
		panic(fmt.Sprintf("Failed to unmarshal zap config: %v", err))
	}

	logger := zap.Must(cfg.Build())
	defer logger.Sync()

	logger.Info("logger construction succeeded")

	return logger
}

func main() {
	r := chi.NewRouter()
	logger := configureLogger()
	userService := services.NewUserService()
	userHandler := handlers.NewUserHandler(logger, userService)

	// Public routes (no auth required)
	r.Post("/users", userHandler.CreateUser)

	// Protected routes (with auth)
	r.Group(func(r chi.Router) {
		r.Use(middlewares.NewAuthMiddleware(logger, userService).BasicAuth)
		r.Get("/users", userHandler.GetUsers)
		r.Get("/users/{id}", userHandler.GetUser)
	})

	err := http.ListenAndServe(":"+constants.DefaultPort, r)
	if err != nil {
		logger.Error("Failed to start server", zap.Error(err))
		return
	}
	logger.Info("Server started", zap.String("port", constants.DefaultPort))
}
