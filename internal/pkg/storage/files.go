package storage

import (
	"fmt"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

func SaveFiles(c *gin.Context, files []*multipart.FileHeader) ([]string, error) {
	var filePaths []string

	for _, file := range files {
		filePath, err := saveFileLocally(c, file)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		filePaths = append(filePaths, filePath)
	}

	return filePaths, nil
}

func saveFileLocally(c *gin.Context, file *multipart.FileHeader) (string, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return "", errors.WithStack(err)
	}

	uploadDir := filepath.Join(cwd, "uploads")

	if err := os.MkdirAll(uploadDir, os.ModePerm); err != nil {
		return "", errors.WithStack(err)
	}

	fileName := fmt.Sprintf("%d%s", time.Now().UnixNano(), filepath.Ext(file.Filename))
	savePath := filepath.Join(uploadDir, fileName)

	fmt.Println("Saving file to path:", savePath)

	if err := c.SaveUploadedFile(file, savePath); err != nil {
		return "", errors.WithStack(err)
	}

	return fmt.Sprintf("http://localhost:8080/media/%s", fileName), nil
}
