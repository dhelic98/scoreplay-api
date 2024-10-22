package service_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/dhelic98/scoreplay-api/application/dto"
	"github.com/dhelic98/scoreplay-api/application/service"
	"github.com/dhelic98/scoreplay-api/domain/entity"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockTagRepository struct {
	mock.Mock
}

func (m *MockTagRepository) CreateTag(ctx context.Context, tag *entity.Tag) error {
	args := m.Called(ctx, tag)
	return args.Error(0)
}

func (m *MockTagRepository) GetAllTags(ctx context.Context) ([]*entity.Tag, error) {
	args := m.Called(ctx)
	return args.Get(0).([]*entity.Tag), args.Error(1)
}

func (m *MockTagRepository) GetTagById(ctx context.Context, id uuid.UUID) (*entity.Tag, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*entity.Tag), args.Error(1)
}

func TestCreateTag(t *testing.T) {
	mockRepo := new(MockTagRepository)
	service := service.NewTagService(mockRepo)

	dto := &dto.CreateTagDTO{
		Name: "Sample Tag",
	}

	mockRepo.On("CreateTag", mock.AnythingOfType("context.backgroundCtx"), mock.AnythingOfType("*entity.Tag")).Return(nil)

	err := service.CreateTag(context.Background(), dto)

	assert.NoError(t, err)

	mockRepo.AssertExpectations(t)
}

func TestGetAllTags(t *testing.T) {
	mockRepo := new(MockTagRepository)
	service := service.NewTagService(mockRepo)

	id1 := uuid.New()
	id2 := uuid.New()

	dtos := []*dto.GetTagDTO{
		{
			ID:   id1,
			Name: "Sample Tag 1",
		}, {
			ID:   id2,
			Name: "Sample Tag 2",
		},
	}

	entities := []*entity.Tag{
		{
			ID:   id1,
			Name: "Sample Tag 1",
		},
		{
			ID:   id2,
			Name: "Sample Tag 2",
		},
	}

	mockRepo.On("GetAllTags", mock.AnythingOfType("context.backgroundCtx")).Return(entities, nil)

	tagsResult, err := service.GetAllTags(context.Background())
	assert.NoError(t, err)

	assert.ElementsMatchf(t, tagsResult, dtos, "Error in DTO array")

	mockRepo.AssertExpectations(t)

}

func TestGetTagById(t *testing.T) {
	mockRepo := new(MockTagRepository)
	service := service.NewTagService(mockRepo)

	tagID, _ := uuid.Parse("123e4567-e89b-12d3-a456-426614174000")

	expectedTag := &dto.GetTagDTO{ID: tagID, Name: "Image1"}
	entity := &entity.Tag{ID: tagID, Name: "Image1"}

	mockRepo.On("GetTagById", mock.AnythingOfType("context.backgroundCtx"), tagID).Return(entity, nil)

	tagResult, err := service.GetTagById(context.Background(), tagID)

	assert.NoError(t, err)
	assert.Equal(t, tagResult, expectedTag)

	mockRepo.AssertExpectations(t)

}

func TestParseMultipartFormToUUID(t *testing.T) {
	mockRepo := new(MockTagRepository)
	service := service.NewTagService(mockRepo)

	id1 := uuid.New()
	id2 := uuid.New()

	uuidArrayString := fmt.Sprintf("[\"%s\",\"%s\"]", id1.String(), id2.String())
	expectedArray := []uuid.UUID{id1, id2}

	result, err := service.ParseMultipartFormToUUID(uuidArrayString)
	assert.NoError(t, err)

	assert.ElementsMatchf(t, result, expectedArray, "UUID arrays not matching")
}
