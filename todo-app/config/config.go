package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppEnv string
	Port   string
	DBURI  string
}

func LoadConfig() *Config {

	// Load .env only for development
	// if os.Getenv("APP_ENV") == "development" {
	// 	if err := godotenv.Load(); err != nil {
	// 		log.Println("No .env file found")

	// 	}
	// }

	_ = godotenv.Load()

	cfg := &Config{
		AppEnv: getEnv("APP_ENV", "development"),
		Port:   getEnv("PORT", "8080"),
		DBURI:  getEnv("DB_URI", ""),
	}

	// Fail fast if critical config missing
	if cfg.DBURI == "" {
		log.Fatal("DB_URI is required but not set")
	}

	return cfg
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
