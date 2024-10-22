package handler_test

import (
	"bytes"
	"context"
	"encoding/json"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/dhelic98/scoreplay-api/application/dto"
	"github.com/dhelic98/scoreplay-api/application/handler"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockImageService struct {
	mock.Mock
}

func (m *MockImageService) CreateImage(ctx context.Context, imageDTO *dto.CreateImageDTO) error {
	args := m.Called(imageDTO)
	return args.Error(0)
}

func (m *MockImageService) GetAllImages(ctx context.Context) ([]*dto.GetImageDTO, error) {
	args := m.Called()
	return args.Get(0).([]*dto.GetImageDTO), args.Error(1)
}

func (m *MockImageService) GetImageById(ctx context.Context, id uuid.UUID) (*dto.GetImageDTO, error) {
	args := m.Called(id)
	return args.Get(0).(*dto.GetImageDTO), args.Error(1)
}

func (m *MockImageService) SearchImagesByTagName(ctx context.Context, tagName string) ([]*dto.GetImageDTO, error) {
	args := m.Called(tagName)
	return args.Get(0).([]*dto.GetImageDTO), args.Error(1)
}

func TestCreateImageHandler(t *testing.T) {
	mockImageService := new(MockImageService)
	mockTagService := new(MockTagService)
	mockFileService := new(MockFileService)

	handler := handler.NewImageHandler(mockImageService, mockTagService, mockFileService)

	tagID := uuid.New()
	expectedImage := dto.CreateImageDTO{
		Name: "Sample Image",
		URL:  "/uploads/test.png",
		Tags: []uuid.UUID{tagID},
	}

	mockImageService.On("CreateImage", &expectedImage).Return(nil)
	mockTagService.On("ParseMultipartFormToUUID", "[\"tag1-uuid\",\"tag2-uuid\"").Return([]uuid.UUID{tagID}, nil)
	mockFileService.On("UploadFile",
		mock.AnythingOfType("*multipart.File"),
		mock.AnythingOfType("*multipart.FileHeader")).Return("/uploads/test.png", nil)

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	_ = writer.WriteField("name", "Sample Image")
	_ = writer.WriteField("tags", "[\"tag1-uuid\",\"tag2-uuid\"")

	fileWriter, err := writer.CreateFormFile("image", "test.png")
	assert.NoError(t, err)

	_, err = fileWriter.Write([]byte("dummy image data"))
	assert.NoError(t, err)

	writer.Close()

	req := httptest.NewRequest(http.MethodPost, "/v1/media", strings.NewReader(body.String()))
	req.Header.Set("Content-Type", writer.FormDataContentType())

	rr := httptest.NewRecorder()

	handler.CreateImageHandler(rr, req)

}

func TestGetAllImagesHandler(t *testing.T) {
	mockImageService := new(MockImageService)
	mockTagService := new(MockTagService)
	mockFileService := new(MockFileService)

	handler := handler.NewImageHandler(mockImageService, mockTagService, mockFileService)

	id1, _ := uuid.Parse("123e4567-e89b-12d3-a456-426614174000")
	id2, _ := uuid.Parse("223e4567-e89b-12d3-a456-426614174001")

	expectedImages := []*dto.GetImageDTO{
		{ID: id1, Name: "Image1", URL: "some url 1"},
		{ID: id2, Name: "Image2", URL: "some url 2"},
	}

	mockImageService.On("GetAllImages").Return(expectedImages, nil)

	req := httptest.NewRequest(http.MethodGet, "/v1/media", nil)

	rr := httptest.NewRecorder()

	handler.GetAllImagesHandler(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var response []map[string]string
	err := json.Unmarshal(rr.Body.Bytes(), &response)
	assert.NoError(t, err)

	assert.Equal(t, "Image1", response[0]["name"])
	assert.Equal(t, "123e4567-e89b-12d3-a456-426614174000", response[0]["id"])
	assert.Equal(t, "some url 1", response[0]["url"])

	assert.Equal(t, "Image2", response[1]["name"])
	assert.Equal(t, "223e4567-e89b-12d3-a456-426614174001", response[1]["id"])
	assert.Equal(t, "some url 2", response[1]["url"])

	mockTagService.AssertExpectations(t)
}

func TestGetImageByIDHandler(t *testing.T) {
	mockImageService := new(MockImageService)
	mockTagService := new(MockTagService)
	mockFileService := new(MockFileService)

	handler := handler.NewImageHandler(mockImageService, mockTagService, mockFileService)

	imageID, _ := uuid.Parse("123e4567-e89b-12d3-a456-426614174000")

	expectedImage := &dto.GetImageDTO{ID: imageID, Name: "Image1", URL: "some url 1"}

	mockImageService.On("GetImageById", imageID).Return(expectedImage, nil)

	req := httptest.NewRequest(http.MethodGet, "/v1/media/1", nil)
	req.SetPathValue("id", imageID.String())

	rr := httptest.NewRecorder()
	handler.GetImageByIDHandler(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var response map[string]string
	err := json.Unmarshal(rr.Body.Bytes(), &response)
	assert.NoError(t, err)

	assert.Equal(t, expectedImage.ID.String(), response["id"])
	assert.Equal(t, expectedImage.Name, response["name"])
	assert.Equal(t, expectedImage.URL, response["url"])

	mockTagService.AssertExpectations(t)
}

func TestSearchByTag(t *testing.T) {
	mockImageService := new(MockImageService)
	mockTagService := new(MockTagService)
	mockFileService := new(MockFileService)

	handler := handler.NewImageHandler(mockImageService, mockTagService, mockFileService)

	id1, _ := uuid.Parse("123e4567-e89b-12d3-a456-426614174000")
	id2, _ := uuid.Parse("223e4567-e89b-12d3-a456-426614174001")

	expectedImages := []*dto.GetImageDTO{
		{ID: id1, Name: "Image1", URL: "some url 1"},
		{ID: id2, Name: "Image2", URL: "some url 2"},
	}

	mockImageService.On("SearchImagesByTagName", "some tag").Return(expectedImages, nil)

	req := httptest.NewRequest(http.MethodGet, "/v1/media/filter/t", nil)
	req.SetPathValue("tagName", "some tag")

	rr := httptest.NewRecorder()
	handler.SearchByTag(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var response []map[string]string
	err := json.Unmarshal(rr.Body.Bytes(), &response)
	assert.NoError(t, err)

	assert.Equal(t, expectedImages[0].ID.String(), response[0]["id"])
	assert.Equal(t, expectedImages[0].Name, response[0]["name"])
	assert.Equal(t, expectedImages[0].URL, response[0]["url"])

	assert.Equal(t, expectedImages[1].ID.String(), response[1]["id"])
	assert.Equal(t, expectedImages[1].Name, response[1]["name"])
	assert.Equal(t, expectedImages[1].URL, response[1]["url"])

	mockTagService.AssertExpectations(t)
}
