package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	APIPort    string
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
}

func LoadConfig() *Config {
	// Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Printf("Warning: Error loading .env file: %v", err)
	}

	return &Config{
		APIPort:    os.Getenv("API_PORT"),
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_PORT"),
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBName:     os.Getenv("DB_NAME"), 
	}
}

func GetEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Printf("WARNING: Environment variable %s is empty", key)
	}
	return value
}
