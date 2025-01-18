package main

import (
	"fmt"
	"net/http"

	"github.com/alae-touba/playing-with-go-chi/config/database"
	"github.com/alae-touba/playing-with-go-chi/config/logger"
	"github.com/alae-touba/playing-with-go-chi/constants"
	"github.com/alae-touba/playing-with-go-chi/handlers"
	"github.com/alae-touba/playing-with-go-chi/middlewares"
	"github.com/alae-touba/playing-with-go-chi/repositories"
	"github.com/alae-touba/playing-with-go-chi/services"
	"github.com/go-chi/chi/v5"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
)

func main() {
	// init logger
	logger, err := logger.ConfigureLogger()
	if err != nil {
		fmt.Printf("Failed to configure logger: %v", err)
		return
	}
	defer logger.Sync()

	if err := database.RunMigrations(); err != nil {
		logger.Fatal("Failed to run migrations", zap.Error(err))
	}
	logger.Info("Migrations ran successfully")

	// init database
	client, err := database.InitDB()
	if err != nil {
		logger.Fatal("failed to initialize database", zap.Error(err))
		return
	}
	defer client.Close()
	logger.Info("Database initialized")

	// handle routes
	r := chi.NewRouter()
	userRepository := repositories.NewUserRepository(client, logger)
	userService := services.NewUserService(logger, userRepository)
	userHandler := handlers.NewUserHandler(logger, userService)
	RegisterRoutes(r, logger, userService, userHandler)

	// Start server
	logger.Info("Server started", zap.String("port", constants.DefaultPort))
	err = http.ListenAndServe(":"+constants.DefaultPort, r)
	if err != nil {
		logger.Error("Failed to start server", zap.Error(err))
		return
	}
}
