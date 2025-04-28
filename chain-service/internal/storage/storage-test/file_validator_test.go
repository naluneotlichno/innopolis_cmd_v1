package storagetest

import (
	"testing"

	"github.com/szaluzhanskaya/Innopolis/chain-service/internal/storage"
)

func TestValidateFile(t *testing.T) {
	tests := []struct {
		name     string
		fileName string
		size     int64
		wantErr  bool
	}{
		{"valid jpg", "test.jpg", 5 << 20, false},
		{"invalid ext", "test.exe", 1 << 20, true},
		{"too large", "test.jpg", 11 << 20, true},
		{"valid mp4", "video.mp4", 1 << 20, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := storage.ValidateFile(tt.fileName, tt.size)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateFile() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
