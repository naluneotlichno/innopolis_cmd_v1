package storage

import (
	"bytes"        // Для создания io.Reader из []byte
	"context"      // Контекст для отмены операций
	"database/sql" // Работа с PostgreSQL
	"fmt"
	"log"
	"path/filepath"
	"strings"

	"github.com/google/uuid"                       // Генерация UUID
	"github.com/minio/minio-go/v7"                 // Клиент MinIO
	"github.com/minio/minio-go/v7/pkg/credentials" // Учетные данные MinIO
)

type MinioConfig struct {
	Endpoint        string // Адрес сервера
	AccessKeyID     string // Логин
	SecretAccessKey string // Пароль
	BucketName      string // Имя кнтейнера для файлов
	Region          string // Регион
	UseSSL          bool   // Использовать SSL
}

// Создание клиета MinIO
func InitMinio(cfg MinioConfig) (*minio.Client, error) {
	return minio.New(cfg.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(cfg.AccessKeyID, cfg.SecretAccessKey, ""), // Указываем ключи
		Secure: cfg.UseSSL,                                                        // Включаем или выключаем SSL
		Region: cfg.Region,                                                        // Указываем регион
	})
}

// Метод загрузки файла в MinIO и записи его в базу данных
func UploadFileMinio(
	ctx context.Context, // Контекст для обработки запросов -- таймауты, отмена, и т. д.
	db *sql.DB, // База данных
	client *minio.Client, // Клиент MinIO
	cfg MinioConfig, // Конфигурация MinIO
	fileName string, // Имя файла
	file []byte, // Файл для загрузки
) (string, string, error) {

	// Проверка расширения файла
	allowedExt := []string{".jpg", ".jpeg", ".png", ".mp4", ".mp3", ".wav"}
	ext := strings.ToLower(filepath.Ext(fileName)) // Получаем расширение файла
	valid := false

	for _, allow := range allowedExt {
		if ext == allow {
			valid = true
			break
		}
	}

	if !valid {
		return "", "", fmt.Errorf("❌ недопустимый тип файла: %s", ext)
	}

	const maxSize = 10 << 20 // 10 MB
	if len(file) > maxSize {
		return "", "", fmt.Errorf("❌ размер файла превышает %d байт", maxSize)
	}

	// Генерация UUID для файла
	storageUUID := uuid.New().String()

	// Генерация пути для хранения файла в MinIO
	objectName := fmt.Sprintf("media/%s%s", storageUUID, ext)

	// Загрузка файла в MinIO
	_, err := client.PutObject(ctx,
		cfg.BucketName,        // Название бакета в MinIO
		objectName,            // Путь к объекту в бакете
		bytes.NewReader(file), // Создание потока из данных
		int64(len(file)),      // Размер файла
		minio.PutObjectOptions{ // Опции, включая Content-Type
			ContentType: "application/octet-stream", // Установите тип контента по умолчанию
		})

	if err != nil {
		return "", "", fmt.Errorf("❌ не удалось загрузить файл в MinIO: %w", err)
	}

	// Сохраняем uuid и путь в базу данных
	_, err = db.ExecContext(ctx, `
		INSERT INTO storage (uuid, s3_path)
		VALUES ($1, $2)
	`, storageUUID, objectName)

	if err != nil {
		return "", "", fmt.Errorf("❌ не удалось записать в БД: %w", err)
	}

	log.Printf("Файл %s успешно загружен как %s", fileName, objectName)

	return storageUUID, objectName, nil
}
