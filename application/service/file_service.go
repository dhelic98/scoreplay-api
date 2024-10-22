package service

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"

	"github.com/google/uuid"
)

type FileService struct {
}

func NewFileService() *FileService {
	return &FileService{}
}

func (f *FileService) UploadFile(file *multipart.File) (string, error) {
	fileID := uuid.New().String()
	filePath := f.GetFilePath(fileID)

	fileBytes, err := io.ReadAll(*file)
	if err != nil {
		return "", err
	}
	err = os.WriteFile(filePath, fileBytes, os.ModePerm)
	if err != nil {
		return "", err
	}

	FILE_HOST_URL := os.Getenv("FILE_HOST_URL")

	fileUrl := fmt.Sprintf("%s/v1/file/%s", FILE_HOST_URL, fileID)

	return fileUrl, nil
}

func (f *FileService) GetFilePath(idStr string) string {
	return fmt.Sprintf("./uploads/%s.jpg", idStr)
}
