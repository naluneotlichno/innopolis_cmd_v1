package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// конфигурация логов
type LogConfig struct {
	Style string `yaml:"style"`
	Level string `yaml:"level"`
	File  string `yaml:"file"`
}

// конфигурация бд
type PostgresConfig struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Port     int    `yaml:"port"`
	Host     string `yaml:"host"`
	DBname   string `yaml:"dbname"`
}

// конфигурация приложения
type AppConfig struct {
	Name        string `yaml:"name"`
	Environment string `yaml:"environment"`
	Port        string `yaml:"port"`
	Debug       bool   `yaml:"debug"`
}

type Config struct {
	Logs      LogConfig      `yaml:"log"`
	DB        PostgresConfig `yaml:"postgres"`
	AppConfig AppConfig      `yaml:"app"`
}

// LoadConfig - функция для загрузки конфигурации из файла
func LoadConfig(env string) (*Config, error) {
	viper.SetConfigName("config." + env) // добавляет к имени значение переменной APP_ENVIRONMENT из .env
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	var cfg Config
	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("ошибка чтения файла конфигурации: %w", err)
	}

	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("ошибка при анмаршалинге файла конфигурации %w", err)
	}

	return &cfg, nil
}
