package handler_test

import (
	"fmt"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/dhelic98/scoreplay-api/application/handler"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockFileService struct {
	mock.Mock
}

func (mock *MockFileService) GetFilePath(fileName string) string {
	args := mock.Called(fileName)
	return args.Get(0).(string)
}

func (mock *MockFileService) UploadFile(file *multipart.File, fileHeader *multipart.FileHeader) (string, error) {
	args := mock.Called(file, fileHeader)
	return args.Get(0).(string), args.Error(1)
}

func TestServeImageFile(t *testing.T) {
	mockFileService := new(MockFileService)
	fileHandler := handler.NewFileHandler(mockFileService)

	fileID := "123e4567-e89b-12d3-a456-426614174000"
	filePath := fmt.Sprintf("http://hostname/v1/file/%s.jpeg", fileID)

	mockFileService.On("GetFilePath", fileID).Return(filePath)

	req := httptest.NewRequest(http.MethodGet, "/v1/file/1", nil)
	req.SetPathValue("fileID", fileID)

	rr := httptest.NewRecorder()
	fileHandler.ServeImageFile(rr, req)
	//Expected status not found since file is not available on server
	assert.Equal(t, http.StatusNotFound, rr.Code)
}
