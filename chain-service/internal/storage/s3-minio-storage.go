package storage

import (
	"bytes"   // Для создания io.Reader из []byte
	"context" // Контекст для отмены операций

	// Работа с PostgreSQL
	"fmt"
	"log"
	"path/filepath"

	"github.com/google/uuid"                       // Генерация UUID
	"github.com/minio/minio-go/v7"                 // Клиент MinIO
	"github.com/minio/minio-go/v7/pkg/credentials" // Учетные данные MinIO
	"github.com/szaluzhanskaya/Innopolis/chain-service/config"
)

// Создание клиента MinIO
func InitMinio(cfg config.MinioConfig) (*minio.Client, error) {
	return minio.New(cfg.Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(cfg.AccessKeyID, cfg.SecretAccessKey, ""), // Указываем ключи
		Secure: cfg.UseSSL,                                                        // Включаем или выключаем SSL
		Region: cfg.Region,                                                        // Указываем регион
	})
}

// Метод загрузки файла в MinIO и записи его в базу данных через repo слой
func UploadFileMinio(
	ctx context.Context, // Контекст для обработки запросов -- таймауты, отмена, и т. д.
	repo StorageRepository, // Репозиторий для работы с таблицей storage
	client *minio.Client, // Клиент MinIO
	cfg config.MinioConfig, // Конфигурация MinIO
	fileName string, // Имя файла
	file []byte, // Файл для загрузки
) (string, string, error) {

	// Проверка расширения файла
	if err := ValidateFile(fileName, int64(len(file))); err != nil {
		return "", "", fmt.Errorf("недопустимый тип файла: %s", err.Error())
	}

	// Генерация UUID для файла
	storageUUID := uuid.New().String()

	// Генерация пути для хранения файла в MinIO
	ext := filepath.Ext(fileName)
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
		return "", "", fmt.Errorf("не удалось загрузить файл в MinIO: %w", err)
	}

	// Сохраняем uuid и путь в базу данных через repo слой
	if err := repo.SaveFileInfo(ctx, storageUUID, objectName); err != nil {
		// Откат: удаляем файл из MinIO
		removeErr := client.RemoveObject(ctx, cfg.BucketName, objectName, minio.RemoveObjectOptions{})
		if removeErr != nil {
			log.Printf("Не удалось удалить файл из MinIO после ошибки БД: %v", removeErr)
		}
		return "", "", err
	}

	log.Printf("Файл %s успешно загружен как %s", fileName, objectName)

	return storageUUID, objectName, nil
}
