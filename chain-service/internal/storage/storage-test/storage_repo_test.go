package storagetest

import (
	"context"
	"testing"

	"go.uber.org/mock/gomock"
)

func TestStorageRepository_SaveFileInfo_WithGoMock(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := NewMockStorageRepository(ctrl)

	// Успешный сценарий
	mockRepo.EXPECT().
		SaveFileInfo(gomock.Any(), "uuid", "path").
		Return(nil)

	err := mockRepo.SaveFileInfo(context.Background(), "uuid", "path")
	if err != nil {
		t.Errorf("ожидалась успешная вставка, но была ошибка: %v", err)
	}

	// Ошибочный сценарий
	mockRepo.EXPECT().
		SaveFileInfo(gomock.Any(), "bad-uuid", "bad-path").
		Return(context.DeadlineExceeded)

	err = mockRepo.SaveFileInfo(context.Background(), "bad-uuid", "bad-path")
	if err == nil {
		t.Error("ожидалась ошибка, но её не было")
	}
}
