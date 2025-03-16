package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Initialize environment variables
func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
}

// GetEnvWithDefault retrieves environment variable or returns default
func GetEnvWithDefault(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
