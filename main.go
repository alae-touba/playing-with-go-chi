package main

import (
	"fmt"
	"net/http"

	"github.com/alae-touba/playing-with-go-chi/config/database"
	"github.com/alae-touba/playing-with-go-chi/config/logger"
	"github.com/alae-touba/playing-with-go-chi/handlers"
	"github.com/alae-touba/playing-with-go-chi/repositories"
	"github.com/alae-touba/playing-with-go-chi/services"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file" // <-- important

	"github.com/alae-touba/playing-with-go-chi/constants"
	"github.com/go-chi/chi/v5"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
)

func main() {
	// Initialize logger
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

	// Initialize database
	client, err := database.InitDB()
	if err != nil {
		logger.Fatal("failed to initialize database", zap.Error(err))
		return
	}
	defer client.Close()
	logger.Info("Database initialized")

	r := chi.NewRouter()

	userRepository := repositories.NewUserRepository(client, logger)
	userService := services.NewUserService(logger, userRepository)
	userHandler := handlers.NewUserHandler(logger, userService)

	// Public routes (no auth required)
	r.Get("/api/v1/hello_world", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello, World!"))
	})

	r.Post("/api/v1/users", userHandler.CreateUser)
	r.Get("/api/v1/users/{id}", userHandler.GetUser)

	// //TODO: refactor
	// r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
	// 	// Test DB connection
	// 	users, err := client.User.Query().All(r.Context())
	// 	if err != nil {
	// 		logger.Error("database query failed", zap.Error(err))
	// 		utils.RespondWithError(w, http.StatusInternalServerError, "database KO")
	// 		return
	// 	}
	// 	logger.Debug("database query succeeded", zap.Int("users_count", len(users)))

	// 	// w.WriteHeader(http.StatusOK)
	// 	utils.RespondWithJSON(w, http.StatusOK, map[string]interface{}{"database": "OK"})
	// 	// json.NewEncoder(w).Encode(map[string]interface{}{
	// 	// 	"status":      "healthy",
	// 	// 	"users_count": len(users),
	// 	// })
	// })

	// // Protected routes (with auth)
	// r.Group(func(r chi.Router) {
	// 	r.Use(middlewares.NewAuthMiddleware(logger, userService).BasicAuth)
	// 	r.Get("/users", userHandler.GetUsers)
	// 	r.Get("/users/{id}", userHandler.GetUser)
	// 	r.Delete("/users/{id}", userHandler.DeleteUser)
	// })

	logger.Info("Server started", zap.String("port", constants.DefaultPort))
	err = http.ListenAndServe(":"+constants.DefaultPort, r)
	if err != nil {
		logger.Error("Failed to start server", zap.Error(err))
		return
	}
}
