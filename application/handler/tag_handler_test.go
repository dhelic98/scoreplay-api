package handler_test

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/dhelic98/scoreplay-api/application/dto"
	"github.com/dhelic98/scoreplay-api/application/handler"
	"github.com/dhelic98/scoreplay-api/domain/entity"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockTagService struct {
	mock.Mock
}

func (mock *MockTagService) CreateTag(c context.Context, dto dto.CreateTagDTO) error {
	args := mock.Called(dto.Name)
	return args.Error(1)
}

func (mock *MockTagService) GetAllTags(ctx context.Context) ([]*dto.GetTagDTO, error) {
	args := mock.Called()
	return args.Get(0).([]*dto.GetTagDTO), args.Error(1)
}

func (mock *MockTagService) GetTagById(ctx context.Context, id uuid.UUID) (*dto.GetTagDTO, error) {
	args := mock.Called(id.String())
	return args.Get(0).(*dto.GetTagDTO), args.Error(1)
}

func (mock *MockTagService) ParseMultipartFormToUUID(tagsJSONString string) ([]uuid.UUID, error) {
	args := mock.Called(tagsJSONString)
	return args.Get(0).([]uuid.UUID), args.Error(1)
}

func TestCreateTag(t *testing.T) {
	mockTagService := new(MockTagService)
	tagHandler := handler.NewTagHandler(mockTagService)

	tagName := "Sample Tag"
	id, _ := uuid.Parse("123e4567-e89b-12d3-a456-426614174000")
	tag := entity.Tag{
		ID:   id,
		Name: tagName,
	}

	mockTagService.On("CreateTag", tagName).Return(tag, nil)

	body, _ := json.Marshal(map[string]string{
		"name": tagName,
	})

	req := httptest.NewRequest(http.MethodPost, "/v1/tags", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()

	tagHandler.CreateTagHandler(rr, req)

	assert.Equal(t, http.StatusCreated, rr.Code)
	mockTagService.AssertExpectations(t)

}

func TestGetAllTagsHandler(t *testing.T) {
	mockTagService := new(MockTagService)
	tagHandler := handler.NewTagHandler(mockTagService)

	id1, _ := uuid.Parse("123e4567-e89b-12d3-a456-426614174000")
	id2, _ := uuid.Parse("223e4567-e89b-12d3-a456-426614174001")
	expectedTags := []*dto.GetTagDTO{
		{ID: id1, Name: "Tag1"},
		{ID: id2, Name: "Tag2"},
	}

	mockTagService.On("GetAllTags").Return(expectedTags, nil)

	req := httptest.NewRequest(http.MethodGet, "/v1/tags", nil)

	rr := httptest.NewRecorder()

	tagHandler.GetAllTagsHandler(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var response []map[string]string
	err := json.Unmarshal(rr.Body.Bytes(), &response)
	assert.NoError(t, err)

	assert.Equal(t, "Tag1", response[0]["name"])
	assert.Equal(t, "123e4567-e89b-12d3-a456-426614174000", response[0]["id"])
	assert.Equal(t, "Tag2", response[1]["name"])
	assert.Equal(t, "223e4567-e89b-12d3-a456-426614174001", response[1]["id"])

	mockTagService.AssertExpectations(t)
}

func TestGetTagByIDHandler(t *testing.T) {
	mockTagService := new(MockTagService)
	tagHandler := handler.NewTagHandler(mockTagService)

	tagIDString := "123e4567-e89b-12d3-a456-426614174000"
	tagID, _ := uuid.Parse(tagIDString)
	expectedTag := &dto.GetTagDTO{
		ID:   tagID,
		Name: "Test Tag",
	}

	mockTagService.On("GetTagById", tagIDString).Return(expectedTag, nil)

	req := httptest.NewRequest(http.MethodGet, "/v1/tags/1", nil)
	req.SetPathValue("id", tagID.String())

	rr := httptest.NewRecorder()
	tagHandler.GetTagByIDHandler(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	var response map[string]string
	err := json.Unmarshal(rr.Body.Bytes(), &response)
	assert.NoError(t, err)

	assert.Equal(t, expectedTag.ID.String(), response["id"])
	assert.Equal(t, expectedTag.Name, response["name"])

	mockTagService.AssertExpectations(t)
}
