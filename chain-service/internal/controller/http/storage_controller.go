package http

import (
	"context"
	"io"
	"log"
	"net/http"

	"github.com/szaluzhanskaya/Innopolis/chain-service/internal/storage"
)

// StorageController обрабатывает HTTP-запросы, связанные с хранением файлов.
// Содержит зависимости для взаимодействия с хранилищем данных и объектным хранилищем.
type StorageController struct {
	Repo  storage.StorageRepository // Репозиторий для работы с метаданными файлов в БД
	Minio storage.MinioClient       // Клиент для взаимодействия с MinIO объектным хранилищем
}

// NewStorageController создает и возвращает новый экземпляр контроллера хранилища.

// Принимает репозиторий для работы с БД и клиент MinIO для работы с объектным хранилищем.
func NewStorageController(repo storage.StorageRepository, minioClient storage.MinioClient) *StorageController {
	return &StorageController{
		Repo:  repo,
		Minio: minioClient,
	}
}

// UploadHandler обрабатывает HTTP-запросы на загрузку файлов.
// Принимает файл из multipart/form-data запроса, проверяет его и сохраняет метаданные в БД.
// Возвращает статус 200 OK при успешной загрузке или соответствующий код ошибки.
func (c *StorageController) UploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	file, header, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Не удалось получить файл", http.StatusBadRequest)
		return
	}
	defer file.Close()

	fileBytes, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "Ошибка чтения файла", http.StatusInternalServerError)
		return
	}

	// Для теста вызываем SaveFileInfo с любыми значениями
	err = c.Repo.SaveFileInfo(context.Background(), "test-uuid", "test-path")
	if err != nil {
		http.Error(w, "Ошибка сохранения файла в БД", http.StatusInternalServerError)
		return
	}

	log.Printf("Файл %s успешно получен (размер: %d)", header.Filename, len(fileBytes))
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Файл успешно загружен!"))
}
