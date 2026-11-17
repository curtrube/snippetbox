package config

import (
	"errors"
	"os"
	"strconv"
)

// Config holds application configuration.
type Config struct {
	Environment string
	MySQL       MySQLConfig
}

// MySQLConfig is MySQL specific database connection configuration.
type MySQLConfig struct {
	Host     string
	Port     int
	Database string
	Username string
	Password string
}

// Load populates the config using environment variables.
func Load() (Config, error) {
	cfg := Config{
		Environment: getEnv("ENVIRONMENT", "dev"),
		MySQL: MySQLConfig{
			Host:     getEnv("MYSQL_HOST", "localhost"),
			Port:     getInt("MYSQL_PORT", 3306),
			Database: getEnv("MYSQL_DATABASE", "snippetbox"),
			Username: getEnv("MYSQL_USERNAME", "web"),
			Password: getEnv("MYSQL_PASSWORD", "pass"),
		},
	}

	if cfg.Environment != "dev" && cfg.MySQL.Password == "" {
		return Config{}, errors.New("MYSQL_PASSWORD is required")
	}

	return cfg, nil
}

// getEnv retrieves the value of the environment variable identified by key.
// If the variable is unset or empty, it returns fallback.
func getEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}

// getInt retrieves the int value of the environment variable identified by key.
// If the variable is unset or empty, it returns fallback.
func getInt(key string, fallback int) int {
	value, err := strconv.Atoi(os.Getenv(key))
	if err != nil {
		return fallback
	}
	return value
}
