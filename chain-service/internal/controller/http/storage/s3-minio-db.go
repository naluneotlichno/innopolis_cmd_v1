package storage

import (
	"context"
	"database/sql"
	"fmt"
)

// SaveFileInfo сохраняет информацию о файле в базу данных
func SaveFileInfo(ctx context.Context, db *sql.DB, uuid, s3Path string) error {
	_, err := db.ExecContext(ctx, `
		INSERT INTO storage (uuid, s3_path)
		VALUES ($1, $2)
	`, uuid, s3Path)

	if err != nil {
		return fmt.Errorf("не удалось записать в БД: %w", err)
	}

	return nil
}
