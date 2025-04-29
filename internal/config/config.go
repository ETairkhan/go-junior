package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadConfig() {
	// Загружаем .env файл
	err := godotenv.Load(".env")
	if err != nil {
		log.Printf("Warning: Error loading .env file: %v", err)
	}

	// Проверяем и выводим значения переменных
	requiredVars := []string{"DB_HOST", "DB_PORT", "DB_USER", "DB_PASSWORD", "DB_NAME", "API_PORT"}
	for _, v := range requiredVars {
		value := os.Getenv(v)
		if value == "" {
			log.Printf("WARNING: Environment variable %s is empty", v)
		} else {
			log.Printf("Environment variable %s is set to: %s", v, value)
		}
	}
}

func GetEnv(key string) string {
	value := os.Getenv(key)
	if value == "" {
		log.Printf("WARNING: Environment variable %s is empty", key)
	}
	return value
}
