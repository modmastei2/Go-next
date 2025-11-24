package config

import (
	"os"

	"github.com/modmastei2/Go-next/backend/pkg/database"
)

// Config holds all application configuration
type Config struct {
	Server   ServerConfig
	Database database.Config
}

// ServerConfig holds server configuration
type ServerConfig struct {
	Port string
	Host string
}

// Load loads configuration from environment variables
func Load() *Config {
	return &Config{
		Server: ServerConfig{
			Port: getEnv("SERVER_PORT", "3001"),
			Host: getEnv("SERVER_HOST", "0.0.0.0"),
		},
		Database: database.Config{
			Driver:   getEnv("DB_DRIVER", "sqlserver"),
			Host:     getEnv("DB_HOST", "localhost"),
			Port:     getEnv("DB_PORT", "1433"),
			User:     getEnv("DB_USER", "sa"),
			Password: getEnv("DB_PASSWORD", "S1u8p3a8#"),
			Database: getEnv("DB_NAME", "ResearchDB"),
		},
	}
}

// getEnv gets an environment variable or returns a default value
func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
