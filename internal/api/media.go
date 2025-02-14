package api

import (
	"fmt"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

const UploadDir = "./uploads/"

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

// сохраняет одиночный файл.
func saveFileLocally(c *gin.Context, file *multipart.FileHeader) (string, error) {
	if err := os.MkdirAll(UploadDir, os.ModePerm); err != nil {
		return "", errors.WithStack(err)
	}

	fileName := fmt.Sprintf("%d%s", time.Now().UnixNano(), filepath.Ext(file.Filename))
	savePath := filepath.Join(UploadDir, fileName)

	if err := c.SaveUploadedFile(file, savePath); err != nil {
		return "", errors.WithStack(err)
	}

	return fmt.Sprintf("http://localhost:8080/media/%s", fileName), nil
}
