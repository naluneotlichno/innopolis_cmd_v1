package config

import (
	"fmt"
	"log"

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

// конфигурация MinIO
type MinioConfig struct {
	Endpoint        string `yaml:"endpoint"`
	AccessKeyID     string `yaml:"accessKeyID"`
	SecretAccessKey string `yaml:"secretAccessKey"`
	BucketName      string `yaml:"bucketName"`
	Region          string `yaml:"region"`
	UseSSL          bool   `yaml:"useSSL"`
}

type Config struct {
	Logs      LogConfig      `yaml:"log"`
	DB        PostgresConfig `yaml:"postgres"`
	AppConfig AppConfig      `yaml:"app"`
	Minio     MinioConfig    `yaml:"minio"`
}

// LoadConfig - функция для загрузки конфигурации из файла
func LoadConfig(env string) (*Config, error) {
	viper.SetConfigName("config." + env) // добавляет к имени значение переменной APP_ENVIRONMENT из .env
	viper.SetConfigType("yaml")
	// Добавляем несколько путей поиска конфигурационного файла
	viper.AddConfigPath(".")        // текущая директория
	viper.AddConfigPath("..")       // родительская директория
	viper.AddConfigPath("../..")    // директория на два уровня выше
	viper.AddConfigPath("../../..") // директория на три уровня выше

	viper.BindEnv("minio.accessKeyID", "MINIO_ACCESS_KEY")
	viper.BindEnv("minio.secretAccessKey", "MINIO_SECRET_ACCESS_KEY")

	var cfg Config
	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("ошибка чтения файла конфигурации: %w", err)
	}

	// Выводим путь к найденному файлу конфигурации для отладки
	log.Printf("Загружен файл конфигурации: %s", viper.ConfigFileUsed())

	// Отладочный вывод ключей конфигурации
	log.Printf("Доступные ключи конфигурации: %v", viper.AllKeys())

	// Проверяем наличие секции postgres
	if viper.IsSet("postgres") {
		log.Printf("Секция postgres найдена в конфигурации")
		log.Printf("postgres.host = %s", viper.GetString("postgres.host"))
		log.Printf("postgres.port = %d", viper.GetInt("postgres.port"))
		log.Printf("postgres.username = %s", viper.GetString("postgres.username"))
		log.Printf("postgres.dbname = %s", viper.GetString("postgres.dbname"))
	} else {
		log.Printf("Секция postgres НЕ найдена в конфигурации")
	}

	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("ошибка при анмаршалинге файла конфигурации %w", err)
	}

	// Прямое присвоение значений из viper
	cfg.DB.Host = viper.GetString("postgres.host")
	cfg.DB.Port = viper.GetInt("postgres.port")
	cfg.DB.Username = viper.GetString("postgres.username")
	cfg.DB.Password = viper.GetString("postgres.password")
	cfg.DB.DBname = viper.GetString("postgres.dbname")

	// Дополнительно устанавливаем параметры app
	cfg.AppConfig.Name = viper.GetString("app.name")
	cfg.AppConfig.Environment = viper.GetString("app.environment")
	cfg.AppConfig.Port = viper.GetString("app.port")
	cfg.AppConfig.Debug = viper.GetBool("app.debug")

	log.Printf("После анмаршалинга: host=%s, port=%d, username=%s, dbname=%s",
		cfg.DB.Host, cfg.DB.Port, cfg.DB.Username, cfg.DB.DBname)
	log.Printf("AppConfig: name=%s, environment=%s, port=%s, debug=%v",
		cfg.AppConfig.Name, cfg.AppConfig.Environment, cfg.AppConfig.Port, cfg.AppConfig.Debug)

	return &cfg, nil
}
