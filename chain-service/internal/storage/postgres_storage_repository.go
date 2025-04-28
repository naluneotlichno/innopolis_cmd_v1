package storage

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

// StorageRepository - интерфейс для работы с таблицей storage
type StorageRepository interface {
	SaveFileInfo(ctx context.Context, uuid, s3Path string) error
}

// PostgresStorageRepository - реализация StorageRepository для PostgreSQL с использованием pgx
type PostgresStorageRepository struct {
	db *pgxpool.Pool
}

// NewPostgresStorageRepository создает новый экземпляр PostgresStorageRepository
func NewPostgresStorageRepository(db *pgxpool.Pool) *PostgresStorageRepository {
	return &PostgresStorageRepository{db: db}
}

// SaveFileInfo сохраняет информацию о файле в базу данных
func (r *PostgresStorageRepository) SaveFileInfo(ctx context.Context, uuid, s3Path string) error {
	_, err := r.db.Exec(ctx, `
		INSERT INTO storage (uuid, s3_path)
		VALUES ($1, $2)
	`, uuid, s3Path)
	if err != nil {
		return fmt.Errorf("не удалось записать в БД: %w", err)
	}
	return nil
}
