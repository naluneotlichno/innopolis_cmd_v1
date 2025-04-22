package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// конфигурация логов
type LogConfig struct {
	Style string `mapstructure:"style"`
	Level string `mapstructure:"level"`
	File  string `mapstructure:"file"`
}

// конфигурация бд
type PostgresConfig struct {
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Port     int    `mapstructure:"port"`
	Host     string `mapstructure:"host"`
	DBname   string `mapstructure:"dbname"`
}

// конфигурация приложения
type AppConfig struct {
	Name        string `mapstructure:"name"`
	Environment string `mapstructure:"environment"`
	Port        string `mapstructure:"port"`
	Debug       bool   `mapstructure:"debug"`
}

type Config struct {
	Logs      LogConfig      `mapstructure:"log"`
	DB        PostgresConfig `mapstructure:"postgres"`
	AppConfig AppConfig      `mapstructure:"app"`
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
