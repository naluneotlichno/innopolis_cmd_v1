package storage

import (
	"context"
	"database/sql"
	"fmt"
)

// StorageRepository определяет интерфейс для работы с таблицей storage
//go:generate mockgen -source=s3-minio-db.go -destination=s3-minio-db_mock.go -package=storage

// StorageRepository - интерфейс для работы с таблицей storage
// (можно мокать для тестов)
type StorageRepository interface {
	SaveFileInfo(ctx context.Context, uuid, s3Path string) error
}

// PostgresStorageRepository - реализация StorageRepository для PostgreSQL
// (db внедряется через конструктор)
type PostgresStorageRepository struct {
	db *sql.DB
}

func NewPostgresStorageRepository(db *sql.DB) *PostgresStorageRepository {
	return &PostgresStorageRepository{db: db}
}

// SaveFileInfo сохраняет информацию о файле в базу данных
func (r *PostgresStorageRepository) SaveFileInfo(ctx context.Context, uuid, s3Path string) error {
	_, err := r.db.ExecContext(ctx, `
		INSERT INTO storage (uuid, s3_path)
		VALUES ($1, $2)
	`, uuid, s3Path)
	if err != nil {
		return fmt.Errorf("не удалось записать в БД: %w", err)
	}
	return nil
}
