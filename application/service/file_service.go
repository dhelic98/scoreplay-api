package service

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"

	"github.com/dhelic98/scoreplay-api/cmd/config"
	"github.com/google/uuid"
)

type IFileService interface {
	UploadFile(file *multipart.File, fileHeader *multipart.FileHeader) (string, error)
	GetFilePath(fileName string) string
}

type FileService struct {
}

func NewFileService() *FileService {
	return &FileService{}
}

func (fileService *FileService) UploadFile(file *multipart.File, fileHeader *multipart.FileHeader) (string, error) {
	fileID := uuid.New().String()
	fileExt := filepath.Ext(fileHeader.Filename)
	fullFileName := fmt.Sprintf("%s%s", fileID, fileExt)
	filePath := fileService.GetFilePath(fullFileName)

	fileBytes, err := io.ReadAll(*file)
	if err != nil {
		return "", err
	}
	err = os.WriteFile(filePath, fileBytes, os.ModePerm)
	if err != nil {
		return "", err
	}

	FILE_HOST_URL := config.GetConfigInstance().FileHostUrl

	fileUrl := fmt.Sprintf("%s/file/%s", FILE_HOST_URL, fullFileName)

	return fileUrl, nil
}

func (fileService *FileService) GetFilePath(fileName string) string {
	return fmt.Sprintf("./uploads/%s", fileName)
}
