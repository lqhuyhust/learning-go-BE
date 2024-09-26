package utils

import (
	"fmt"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
)

func UploadFile(c *gin.Context, formField string, uploadDir string) (string, error) {
	file, err := c.FormFile(formField)
	if err != nil {
		return "", err
	}

	// make file name
	timeStamp := time.Now().Unix()
	originalFileName := filepath.Base(file.Filename)
	uniqueFileName := fmt.Sprintf("%d_%s", timeStamp, originalFileName)
	profilePath := filepath.Join(uploadDir, uniqueFileName)

	// save file
	if err := c.SaveUploadedFile(file, profilePath); err != nil {
		return "", err
	}

	return profilePath, nil
}
