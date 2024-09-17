package config

import (
	"os"
)

// DBConfig struct holds the database credentials
type DBConfig struct {
	User     string
	Password string
	DBName   string
	Host     string
	Port     string
	SSLMode  string
}

// GetDBConfig returns the database configuration
func GetDBConfig() *DBConfig {
	return &DBConfig{
		User:     getEnv("DB_USER", "postgres"),
		Password: getEnv("DB_PASSWORD", "password"),
		DBName:   getEnv("DB_NAME", "mydb"),
		Host:     getEnv("DB_HOST", "localhost"),
		Port:     getEnv("DB_PORT", "5432"),
		SSLMode:  getEnv("DB_SSLMODE", "disable"), // disable for local development
	}
}

// Helper function to get environment variables or return a default value
func getEnv(key string, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
