package database

import (
	"context"
	"fmt"
	"os"
	"time"

	"entgo.io/ent/dialect"
	"github.com/alae-touba/playing-with-go-chi/repositories/ent"
	"github.com/golang-migrate/migrate/v4"
)

func InitDB() (*ent.Client, error) {
	dbUser := getEnvOrDefault("DB_USER", "postrgres")
	dbPassword := getEnvOrDefault("DB_PASSWORD", "postrgres")
	dbName := getEnvOrDefault("DB_NAME", "postrgres")
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

func RunMigrations() error {
	dbUser := getEnvOrDefault("DB_USER", "postrgres")
	dbPassword := getEnvOrDefault("DB_PASSWORD", "postrgres")
	dbName := getEnvOrDefault("DB_NAME", "postrgres")
	dbPort := getEnvOrDefault("DB_PORT", "5432")

	m, err := migrate.New(
		"file://migrations",
		fmt.Sprintf("postgres://%s:%s@postgres:%s/%s?sslmode=disable",
			dbUser, dbPassword, dbPort, dbName),
	)
	if err != nil {
		return err
	}
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return err
	}
	return nil
}

func getEnvOrDefault(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	return value
}
