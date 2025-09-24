package config

import (
	"bileygr/models"
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"
)

func getEnv(key string, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func Load() (*models.Config, error) {
	godotenv.Load("../.env")

	cfg := &models.Config{}

	// Server
	cfg.Server.Port = getEnv("SERVER_PORT", "8080")
	cfg.Server.Host = getEnv("SERVER_HOST", "0.0.0.0")
	cfg.Server.ReadTimeout = time.Second * 15
	cfg.Server.WriteTimeout = time.Second * 15

	// Database
	cfg.Database.Host = getEnv("DB_HOST", "localhost")
	cfg.Database.Port = getEnv("DB_PORT", "5432")
	cfg.Database.User = getEnv("DB_USER", "harrysharma")
	cfg.Database.Password = getEnv("DB_PASSWORD", "")
	cfg.Database.DBName = getEnv("DB_NAME", "bileygr")
	cfg.Database.SSLMode = getEnv("DB_SSLMODE", "disable")

	// JWT config
	cfg.JWT.Secret = getEnv("JWT_SECRET", "your-secret-key")
	cfg.JWT.TokenExpiry = time.Hour * 24    // 24 hours
	cfg.JWT.RefreshExpiry = time.Hour * 168 // 7 days

	cfg.Environment = getEnv("ENV", "development")

	return cfg, nil
}

func GetDSN(cfg *models.Config) string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.DBName,
		cfg.Database.SSLMode,
	)
}
