package config

import (
	"os"
)

type LogConfig struct {
	Style string
	Level string
}

type PostgresConfig struct {
	Username string
	Password string
	Port     string
}

type Config struct {
	Logs LogConfig
	DB   PostgresConfig
	Port string
}

// LoadConfig - функция для загрузки конфигурации из файла
func LoadConfig() (*Config, error) {
	cfg := &Config{
		Port: os.Getenv("PORT"),
		Logs: LogConfig{
			Style: os.Getenv("LOG_STYLE"), // содержит стиль вывода логов
			Level: os.Getenv("LOG_LEVEL"), // содержит тэг классификации логов
		},
		DB: PostgresConfig{
			Username: os.Getenv("PG_USER"),
			Password: os.Getenv("PG_PASSWORD"),
			Port:     os.Getenv("PG_PORT"),
		},
	}
	return cfg, nil
}
