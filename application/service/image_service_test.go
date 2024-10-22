package service_test

import (
	"context"
	"testing"

	"github.com/dhelic98/scoreplay-api/application/dto"
	"github.com/dhelic98/scoreplay-api/application/service"
	"github.com/dhelic98/scoreplay-api/domain/entity"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockImageRepository struct {
	mock.Mock
}

func (m *MockImageRepository) CreateImage(ctx context.Context, image *entity.Image) error {
	args := m.Called(ctx, image)
	return args.Error(0)
}

func (m *MockImageRepository) GetAllImages(ctx context.Context) ([]*entity.Image, error) {
	args := m.Called(ctx)
	return args.Get(0).([]*entity.Image), args.Error(1)
}

func (m *MockImageRepository) GetImageByID(ctx context.Context, id uuid.UUID) (*entity.Image, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*entity.Image), args.Error(1)
}

func (m *MockImageRepository) SearchByTagName(ctx context.Context, tagName string) ([]*entity.Image, error) {
	args := m.Called(ctx, tagName)
	return args.Get(0).([]*entity.Image), args.Error(1)
}

func TestCreateImage(t *testing.T) {
	mockRepo := new(MockImageRepository)
	service := service.NewImageService(mockRepo)

	id1 := uuid.New()
	id2 := uuid.New()

	dto := &dto.CreateImageDTO{
		Name: "Sample Image",
		Tags: []uuid.UUID{id1, id2},
		URL:  "/uploads/sample.jpg",
	}

	mockRepo.On("CreateImage", mock.AnythingOfType("context.backgroundCtx"), mock.AnythingOfType("*entity.Image")).Return(nil)

	err := service.CreateImage(context.Background(), dto)

	assert.NoError(t, err)

	mockRepo.AssertExpectations(t)
}

func TestGetAllImages(t *testing.T) {
	mockRepo := new(MockImageRepository)
	service := service.NewImageService(mockRepo)

	id1 := uuid.New()
	id2 := uuid.New()

	dtos := []*dto.GetImageDTO{
		{
			ID:   id1,
			Name: "Sample Image 1",
			Tags: []string{"test 1 1", "test 1 2"},
			URL:  "/uploads/sample.jpg",
		}, {
			ID:   id2,
			Name: "Sample Image 2",
			Tags: []string{"test 2 1", "test 2 2"},
			URL:  "/uploads/sample.jpg",
		},
	}

	entities := []*entity.Image{
		{
			ID:   id1,
			Name: "Sample Image 1",
			Tags: []entity.Tag{
				{
					ID:   id1,
					Name: "test 1 1",
				},
				{
					ID:   id2,
					Name: "test 1 2",
				},
			},
			URL: "/uploads/sample.jpg",
		},
		{
			ID:   id2,
			Name: "Sample Image 2",
			Tags: []entity.Tag{
				{
					ID:   id1,
					Name: "test 2 1",
				},
				{
					ID:   id2,
					Name: "test 2 2",
				},
			},
			URL: "/uploads/sample.jpg",
		},
	}

	mockRepo.On("GetAllImages", mock.AnythingOfType("context.backgroundCtx")).Return(entities, nil)

	imagesResult, err := service.GetAllImages(context.Background())
	assert.NoError(t, err)

	assert.ElementsMatchf(t, imagesResult, dtos, "Error in DTO array")

	mockRepo.AssertExpectations(t)
}

func TestGetImageById(t *testing.T) {
	mockRepo := new(MockImageRepository)
	service := service.NewImageService(mockRepo)

	imageID, _ := uuid.Parse("123e4567-e89b-12d3-a456-426614174000")

	expectedImage := &dto.GetImageDTO{ID: imageID, Name: "Image1", URL: "some url 1", Tags: []string{"tag1", "tag2"}}
	id1 := uuid.New()
	id2 := uuid.New()
	entity := &entity.Image{
		ID: imageID, Name: "Image1", URL: "some url 1", Tags: []entity.Tag{
			{
				ID:   id1,
				Name: "tag1",
			},
			{
				ID:   id2,
				Name: "tag2",
			},
		}}

	mockRepo.On("GetImageByID", mock.AnythingOfType("context.backgroundCtx"), imageID).Return(entity, nil)

	imageResult, err := service.GetImageById(context.Background(), imageID)

	assert.NoError(t, err)
	assert.Equal(t, imageResult, expectedImage)

	mockRepo.AssertExpectations(t)

}

func TestSearchImagesByTagName(t *testing.T) {
	mockRepo := new(MockImageRepository)
	service := service.NewImageService(mockRepo)

	id1 := uuid.New()
	id2 := uuid.New()

	dtos := []*dto.GetImageDTO{
		{
			ID:   id1,
			Name: "Sample Image 1",
			Tags: []string{"tag1", "test 1 2"},
			URL:  "/uploads/sample.jpg",
		}, {
			ID:   id2,
			Name: "Sample Image 2",
			Tags: []string{"tag1", "test 2 2"},
			URL:  "/uploads/sample.jpg",
		},
	}

	entities := []*entity.Image{
		{
			ID:   id1,
			Name: "Sample Image 1",
			Tags: []entity.Tag{
				{
					ID:   id1,
					Name: "tag1",
				},
				{
					ID:   id2,
					Name: "test 1 2",
				},
			},
			URL: "/uploads/sample.jpg",
		},
		{
			ID:   id2,
			Name: "Sample Image 2",
			Tags: []entity.Tag{
				{
					ID:   id1,
					Name: "tag1",
				},
				{
					ID:   id2,
					Name: "test 2 2",
				},
			},
			URL: "/uploads/sample.jpg",
		},
	}

	mockRepo.On("SearchByTagName", mock.AnythingOfType("context.backgroundCtx"), "tag 1").Return(entities, nil)

	imagesResult, err := service.SearchImagesByTagName(context.Background(), "tag 1")
	assert.NoError(t, err)

	assert.ElementsMatchf(t, imagesResult, dtos, "Error in DTO array")

	mockRepo.AssertExpectations(t)
}
