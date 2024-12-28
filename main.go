package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"context"

	"entgo.io/ent/dialect"
	"github.com/alae-touba/playing-with-go-chi/constants"
	"github.com/alae-touba/playing-with-go-chi/handlers"
	"github.com/alae-touba/playing-with-go-chi/middlewares"
	"github.com/alae-touba/playing-with-go-chi/repositories"
	"github.com/alae-touba/playing-with-go-chi/repositories/ent"
	"github.com/alae-touba/playing-with-go-chi/services"
	"github.com/alae-touba/playing-with-go-chi/utils"
	"github.com/go-chi/chi/v5"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
)

func configureLogger() (*zap.Logger, error) {
	configFile, err := os.ReadFile("config/zap_config.json")
	if err != nil {
		return nil, fmt.Errorf("Failed to read zap config file: %v", err)
	}

	var cfg zap.Config
	if err := json.Unmarshal(configFile, &cfg); err != nil {
		return nil, fmt.Errorf("Failed to unmarshal zap config: %v", err)
	}

	logger, err := cfg.Build()
	if err != nil {
		return nil, fmt.Errorf("Failed to build zap logger: %v", err)
	}

	logger.Info("logger construction succeeded")

	return logger, nil
}

func initDB() (*ent.Client, error) {
	dbUser := getEnvOrDefault("DB_USER", "postgres")
	dbPassword := getEnvOrDefault("DB_PASSWORD", "postgres")
	dbName := getEnvOrDefault("DB_NAME", "postgres")
	dbPort := getEnvOrDefault("DB_PORT", "5432")

	dsn := fmt.Sprintf("host=postgres user=%s password=%s dbname=%s port=%s sslmode=disable",
		dbUser, dbPassword, dbName, dbPort)

	// Add retry logic
	var client *ent.Client
	var err error

	for i := 0; i < 5; i++ {
		client, err = ent.Open(dialect.Postgres, dsn)
		if err == nil {
			break
		}
		time.Sleep(time.Second * 2)
	}
	if err != nil {
		return nil, fmt.Errorf("failed opening connection to postgres: %v", err)
	}

	// Verify connection
	ctx := context.Background()
	if err := client.Schema.Create(ctx); err != nil {
		return nil, fmt.Errorf("failed creating schema resources: %v", err)
	}

	return client, nil
}

func getEnvOrDefault(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	return value
}

func main() {
	// Initialize logger
	logger, err := configureLogger()
	if err != nil {
		fmt.Printf("Failed to configure logger: %v", err)
		return
	}
	defer logger.Sync()

	// Initialize database
	client, err := initDB()
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
	r.Post("/users", userHandler.CreateUser)

	//TODO: refactor
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		// Test DB connection
		users, err := client.User.Query().All(r.Context())
		if err != nil {
			logger.Error("database query failed", zap.Error(err))
			utils.RespondWithError(w, http.StatusInternalServerError, "database KO")
			return
		}
		logger.Debug("database query succeeded", zap.Int("users_count", len(users)))

		// w.WriteHeader(http.StatusOK)
		utils.RespondWithJSON(w, http.StatusOK, map[string]interface{}{"database": "OK"})
		// json.NewEncoder(w).Encode(map[string]interface{}{
		// 	"status":      "healthy",
		// 	"users_count": len(users),
		// })
	})

	// Protected routes (with auth)
	r.Group(func(r chi.Router) {
		r.Use(middlewares.NewAuthMiddleware(logger, userService).BasicAuth)
		r.Get("/users", userHandler.GetUsers)
		r.Get("/users/{id}", userHandler.GetUser)
	})

	err = http.ListenAndServe(":"+constants.DefaultPort, r)
	if err != nil {
		logger.Error("Failed to start server", zap.Error(err))
		return
	}
	logger.Info("Server started", zap.String("port", constants.DefaultPort))
}
