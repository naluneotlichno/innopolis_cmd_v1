package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"

	//"github.com/go-delve/delve/pkg/config"
	"github.com/joho/godotenv"
	"github.com/szaluzhanskaya/Innopolis/chain-service/config"
	v1 "github.com/szaluzhanskaya/Innopolis/chain-service/internal/controller/http"
	"github.com/szaluzhanskaya/Innopolis/chain-service/internal/storage"
)

// main - точка входа в приложение chain-service.
// Сервис предоставляет API для загрузки и хранения файлов с использованием
// PostgreSQL для метаданных и MinIO для хранения самих файлов.
func main() {
	// Вывод текущей директории для отладки
	dir, _ := os.Getwd()
	log.Printf("Текущая директория: %s", dir)

	// Загрузка переменных окружения из .env файла в текущей директории
	if err := godotenv.Load(".env"); err != nil {
		log.Printf("Предупреждение: не удалось загрузить .env файл: %v", err)
	}

	// Определение окружения запуска приложения (local, dev, prod)
	// По умолчанию используется local, если переменная не задана
	env := os.Getenv("APP_ENVIRONMENT")
	if env == "" {
		env = "local"
		log.Printf("APP_ENVIRONMENT не задан, используется значение по умолчанию: %s", env)
	}

	// Загрузка конфигурации в зависимости от окружения
	// Конфигурация содержит настройки для БД, MinIO и самого приложения
	log.Printf("Загрузка конфигурации для окружения: %s", env)
	cfg, err := config.LoadConfig(env)
	if err != nil {
		log.Fatalf("Критическая ошибка при загрузке конфигурации: %v", err)
	}
	log.Printf("Конфигурация успешно загружена")

	// Проверка параметров подключения к БД
	log.Printf("Параметры подключения к PostgreSQL: host=%s, port=%d, username=%s, dbname=%s",
		cfg.DB.Host, cfg.DB.Port, cfg.DB.Username, cfg.DB.DBname)

	// Инициализация соединения с PostgreSQL
	// Используется для хранения метаданных о загруженных файлах
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.DB.Host,
		cfg.DB.Port,
		cfg.DB.Username,
		cfg.DB.Password,
		cfg.DB.DBname,
	)
	log.Printf("Строка подключения к PostgreSQL: %s", dsn)

	var db *sql.DB
	var repo storage.StorageRepository

	db, err = sql.Open("postgres", dsn)
	if err != nil {
		log.Printf("ОШИБКА: Не удалось инициализировать подключение к PostgreSQL: %v", err)
		log.Printf("ВНИМАНИЕ: Продолжаем без PostgreSQL, некоторые функции будут недоступны")
		// Создаем заглушку для репозитория
		repo = &storage.MockStorageRepository{}
	} else {
		// Проверяем соединение
		err = db.Ping()
		if err != nil {
			log.Printf("ОШИБКА: PostgreSQL недоступен: %v", err)
			log.Printf("Проверьте, что сервер PostgreSQL запущен и доступен по адресу %s:%d", cfg.DB.Host, cfg.DB.Port)
			log.Printf("ВНИМАНИЕ: Продолжаем без PostgreSQL, некоторые функции будут недоступны")
			db.Close()
			// Создаем заглушку для репозитория
			repo = &storage.MockStorageRepository{}
		} else {
			log.Println("Успешное подключение к PostgreSQL")
			// Создание репозитория для работы с базой данных
			repo = storage.NewPostgresStorageRepository(db)
			defer db.Close()
		}
	}

	// Инициализация клиента MinIO для хранения файлов
	var minioClient storage.MinioClient
	tmpMinioClient, err := storage.InitMinio(cfg.Minio)
	if err != nil {
		log.Printf("ОШИБКА: Не удалось инициализировать клиент MinIO: %v", err)
		log.Printf("ВНИМАНИЕ: Продолжаем без MinIO, некоторые функции будут недоступны")
		// Создаем заглушку для MinIO клиента
		minioClient = &storage.MockMinioClient{}
	} else {
		log.Println("Успешное подключение к MinIO")
		// Оборачиваем клиент MinIO в наш адаптер
		minioClient = storage.NewMinioClientAdapter(tmpMinioClient)
	}

	// Создание контроллера для обработки HTTP-запросов, связанных с хранилищем
	storageController := v1.NewStorageController(repo, minioClient)

	// Регистрация HTTP-обработчиков
	// /ping - для проверки доступности сервиса
	http.HandleFunc("/ping", v1.PingHandler)

	// /upload - для загрузки файлов в хранилище
	http.HandleFunc("/upload", storageController.UploadHandler)

	// Запуск HTTP-сервера на порту, указанном в конфигурации
	serverAddr := ":" + cfg.AppConfig.Port
	log.Printf("Запуск HTTP-сервера на порту %s (адрес %s)...", cfg.AppConfig.Port, serverAddr)
	log.Fatal(http.ListenAndServe(serverAddr, nil))
}
