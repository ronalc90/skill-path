package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// Config holds all configuration values for the application.
type Config struct {
	ServerPort   string
	GinMode      string
	DBPath       string
	JWTSecret    string
	JWTExpiryHrs int
	CORSOrigins  string
}

// Load reads configuration from environment variables.
func Load() (*Config, error) {
	_ = godotenv.Load()

	jwtExpiry, _ := strconv.Atoi(getEnv("JWT_EXPIRY_HOURS", "72"))

	cfg := &Config{
		ServerPort:   getEnv("SERVER_PORT", "3005"),
		GinMode:      getEnv("GIN_MODE", "debug"),
		DBPath:       getEnv("DB_PATH", "skillpath.db"),
		JWTSecret:    getEnv("JWT_SECRET", "default-secret-change-me"),
		JWTExpiryHrs: jwtExpiry,
		CORSOrigins:  getEnv("CORS_ORIGINS", "http://localhost:3000,http://localhost:3015"),
	}

	return cfg, nil
}

func getEnv(key, fallback string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	return fallback
}
