package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

func LoadAppConfig() (AppConfig, error) {
	err := godotenv.Load()
	if err != nil {
		return AppConfig{}, err
	}

	appConfig := AppConfig{
		Database: DatabaseConfig{
			Host:     os.Getenv("DB_HOST"),
			Port:     getIntEnv("DB_PORT"),
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			Name:     os.Getenv("DB_NAME"),
		},
		Server: APIConfig{
			Port: getIntEnv("SERVER_PORT"),
		},
		Webhook: WebhookConfig{
			BaseURL: os.Getenv("WEBHOOK_URL"),
			Path:    os.Getenv("WEBHOOK_PATH"),
		},
	}

	fmt.Println(appConfig)
	return appConfig, nil
}

func getIntEnv(key string) int {
	valStr := os.Getenv(key)
	val, _ := strconv.Atoi(valStr) // Assume error handling is managed elsewhere
	return val
}
