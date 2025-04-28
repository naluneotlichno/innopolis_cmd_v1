package storage

import (
	"context"
	"io"
	"os"
	"path/filepath"
	"time"

	"github.com/google/uuid"
)

type StorageService struct {
	basePath string
}

func NewStorageService(basePath string) *StorageService {
	return &StorageService{
		basePath: basePath,
	}
}

func (s *StorageService) UploadFile(ctx context.Context, file io.Reader, filename string) (string, error) {
	// Генерируем уникальный ID для файла
	fileID := uuid.New().String() + filepath.Ext(filename)

	// Создаем путь для сохранения файла
	filePath := filepath.Join(s.basePath, fileID)

	// Создаем директорию, если она не существует
	if err := os.MkdirAll(s.basePath, 0755); err != nil {
		return "", err
	}

	// Создаем файл
	dst, err := os.Create(filePath)
	if err != nil {
		return "", err
	}
	defer dst.Close()

	// Копируем содержимое файла
	if _, err := io.Copy(dst, file); err != nil {
		return "", err
	}

	return fileID, nil
}

func (s *StorageService) GetFile(ctx context.Context, fileID string) (*os.File, error) {
	filePath := filepath.Join(s.basePath, fileID)

	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	return file, nil
}

func (s *StorageService) CleanupOldFiles(ctx context.Context, maxAge time.Duration) error {
	files, err := os.ReadDir(s.basePath)
	if err != nil {
		return err
	}

	now := time.Now()
	for _, file := range files {
		info, err := file.Info()
		if err != nil {
			continue
		}

		if now.Sub(info.ModTime()) > maxAge {
			os.Remove(filepath.Join(s.basePath, file.Name()))
		}
	}

	return nil
}
