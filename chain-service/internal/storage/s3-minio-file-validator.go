package storage

import (
	"fmt"
	"path/filepath"
	"strings"
)

const maxSize = 10 << 20 // 10 MB

// ValidateFile проверяет файл на соответствие требованиям
func ValidateFile(fileName string, fileSize int64) error {
	// Проверка расширения файла
	allowedExt := []string{".jpg", ".jpeg", ".png", ".mp4", ".mp3", ".wav"}
	ext := strings.ToLower(filepath.Ext(fileName))
	valid := false

	for _, allow := range allowedExt {
		if ext == allow {
			valid = true
			break
		}
	}

	if !valid {
		return fmt.Errorf("недопустимый тип файла: %s", ext)
	}

	if fileSize > maxSize {
		return fmt.Errorf("размер файла превышает %d байт", maxSize)
	}

	return nil
}
