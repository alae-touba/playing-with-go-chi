package logger

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"go.uber.org/zap"
)

func ConfigureLogger() (*zap.Logger, error) {
	configPath := filepath.Join("config", "logger", "zap_config.json")
	configFile, err := os.ReadFile(configPath)
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
