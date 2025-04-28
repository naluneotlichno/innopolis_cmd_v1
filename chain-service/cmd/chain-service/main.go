package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"github.com/szaluzhanskaya/Innopolis/chain-service/config"
	v1 "github.com/szaluzhanskaya/Innopolis/chain-service/internal/controller/http"
	"github.com/szaluzhanskaya/Innopolis/chain-service/internal/repo/postgres"
	"github.com/szaluzhanskaya/Innopolis/chain-service/internal/storage"
	"github.com/szaluzhanskaya/Innopolis/chain-service/internal/usecase"
)

// main - точка входа в приложение chain-service.
// Сервис предоставляет API для загрузки и хранения файлов с использованием
// PostgreSQL для метаданных и MinIO для хранения самих файлов.
func main() {
	// Вывод текущей директории для отладки
	dir, _ := os.Getwd()
	log.Printf("Текущая директория: %s", dir)

	// Загрузка переменных окружения из .env файла
	if err := godotenv.Load(".env"); err != nil {
		log.Printf("Предупреждение: не удалось загрузить .env файл: %v", err)
	}

	// Определение окружения запуска приложения
	env := os.Getenv("APP_ENVIRONMENT")
	if env == "" {
		env = "local"
		log.Printf("APP_ENVIRONMENT не задан, используется значение по умолчанию: %s", env)
	}

	// Загрузка конфигурации
	log.Printf("Загрузка конфигурации для окружения: %s", env)
	cfg, err := config.LoadConfig(env)
	if err != nil {
		log.Fatalf("Критическая ошибка при загрузке конфигурации: %v", err)
	}
	log.Printf("Конфигурация успешно загружена")

	// Инициализация соединения с PostgreSQL
	connURL := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		cfg.DB.Username,
		cfg.DB.Password,
		cfg.DB.Host,
		cfg.DB.Port,
		cfg.DB.DBname,
	)

	dbPool, err := initDb(connURL)
	if err != nil {
		log.Fatalf("Ошибка подключения к базе данных: %v", err)
	}
	defer dbPool.Close()

	// Инициализация репозиториев
	messageChainRepo := postgres.New(dbPool)
	storageRepo := storage.NewPostgresStorageRepository(dbPool)

	// Инициализация клиента MinIO для хранения файлов
	var minioClient storage.MinioClient
	tmpMinioClient, err := storage.InitMinio(cfg.Minio)
	if err != nil {
		log.Printf("ОШИБКА: Не удалось инициализировать клиент MinIO: %v", err)
		log.Printf("ВНИМАНИЕ: Продолжаем без MinIO, некоторые функции будут недоступны")
		minioClient = &storage.MockMinioClient{}
	} else {
		log.Println("Успешное подключение к MinIO")
		minioClient = storage.NewMinioClientAdapter(tmpMinioClient)
	}

	// Создание сервисов и обработчиков
	service := usecase.New(messageChainRepo)
	handler := v1.New(service)
	storageController := v1.NewStorageController(storageRepo, minioClient)

	// Регистрация HTTP-обработчиков
	http.HandleFunc("/ping", v1.PingHandler)
	http.HandleFunc("/delete-chain/{uuid}", handler.DeleteMessageChain)
	http.HandleFunc("/create-chain", handler.CreateMessageChain)
	http.HandleFunc("/upload", storageController.UploadHandler)

	// Запуск HTTP-сервера
	serverAddr := ":" + cfg.AppConfig.Port
	log.Printf("Запуск HTTP-сервера на порту %s (адрес %s)...", cfg.AppConfig.Port, serverAddr)
	log.Fatal(http.ListenAndServe(serverAddr, nil))
}

func initDb(connURL string) (*pgxpool.Pool, error) {
	dbPool, err := pgxpool.New(context.Background(), connURL)
	if err != nil {
		return nil, err
	}
	return dbPool, nil
}
