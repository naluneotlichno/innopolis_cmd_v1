package http

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/szaluzhanskaya/Innopolis/chain-service/internal/storage"
	storagetest "github.com/szaluzhanskaya/Innopolis/chain-service/internal/storage/storage-test"
	"go.uber.org/mock/gomock"
)

func TestUploadHandler_Success(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockRepo := storagetest.NewMockStorageRepository(mockCtrl)
	mockRepo.EXPECT().SaveFileInfo(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)

	controller := &StorageController{
		Repo:  mockRepo,
		Minio: &storage.MockMinioClient{},
	}

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, _ := writer.CreateFormFile("file", "test.txt")
	io.Copy(part, bytes.NewBufferString("hello world"))
	writer.Close()

	req := httptest.NewRequest(http.MethodPost, "/upload", body)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	w := httptest.NewRecorder()
	controller.UploadHandler(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Файл успешно загружен!")
}
