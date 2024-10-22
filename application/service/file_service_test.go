package service_test

import (
	"testing"

	"github.com/dhelic98/scoreplay-api/application/service"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockFileRepository struct {
	mock.Mock
}

func TestGetFilePath(t *testing.T) {
	service := service.NewFileService()

	fileName := "someImage.jpg"
	expectedFilePath := "./uploads/someImage.jpg"

	filePath := service.GetFilePath(fileName)

	assert.Equal(t, expectedFilePath, filePath)

}
