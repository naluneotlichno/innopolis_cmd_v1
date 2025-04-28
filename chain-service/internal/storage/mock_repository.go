package storage

import (
	"context"
)

// MockStorageRepository - заглушка для StorageRepository
// Нужна только для запуска main.go без PostgreSQL
type MockStorageRepository struct{}

// SaveFileInfo - заглушка для метода сохранения информации о файле
func (m *MockStorageRepository) SaveFileInfo(ctx context.Context, uuid, s3Path string) error {
	return nil
}

// MinioClient - интерфейс для MinIO
// Нужен только для запуска main.go
type MinioClient interface {
	// Пустой интерфейс - нам важно чтобы только компилировалось
}

// MockMinioClient - заглушка для MinIO
type MockMinioClient struct{}

// Получаем клиент Minio как MinioClient для контроллера
func NewMinioClientAdapter(client interface{}) MinioClient {
	// Просто возвращаем любой объект как MinioClient
	return &MockMinioClient{}
}
