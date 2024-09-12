package utils

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"
)

func HandleFileUpload(file multipart.File, header *multipart.FileHeader) (string, error) {
	originalFileName := header.Filename
	uniqueFileName := fmt.Sprintf("%d_%s", time.Now().Unix(), originalFileName)
	profilePath := filepath.Join("uploads", uniqueFileName)

	f, err := os.Create(profilePath)
	if err != nil {
		return "", fmt.Errorf("failed to create file: %v", err)
	}
	defer f.Close()

	_, err = io.Copy(f, file)
	if err != nil {
		return "", fmt.Errorf("failed to write file: %v", err)
	}

	return profilePath, nil
}
