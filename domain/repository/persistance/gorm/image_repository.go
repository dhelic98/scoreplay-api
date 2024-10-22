package gorm

import (
	"context"

	"github.com/dhelic98/scoreplay-api/domain/entity"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ImageRepository struct {
	DB *gorm.DB
}

func NewPostgresSQLImageRepository(db *gorm.DB) *ImageRepository {
	return &ImageRepository{DB: db}
}

// Write
func (repository *ImageRepository) CreateImage(ctx context.Context, image *entity.Image) error {
	return repository.DB.Create(image).Error
}

// Read
func (repository *ImageRepository) GetImageByID(ctx context.Context, id uuid.UUID) (*entity.Image, error) {
	var image entity.Image
	if err := repository.DB.Preload("Tags").First(&image, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &image, nil
}

func (repository *ImageRepository) GetAllImages(ctx context.Context) ([]*entity.Image, error) {
	var images []*entity.Image
	if err := repository.DB.Preload("Tags").Find(&images).Error; err != nil {
		return nil, err
	}
	return images, nil
}

func (repository *ImageRepository) SearchByTagName(ctx context.Context, tagName string) ([]*entity.Image, error) {
	var images []*entity.Image
	if err := repository.DB.Joins("JOIN image_tags ON images.id = image_tags.image_id").
		Joins("JOIN tags ON image_tags.tag_id = tags.id").
		Where("tags.name = ?", tagName).
		Preload("Tags").
		Find(&images).Error; err != nil {
		return nil, err
	}

	return images, nil
}
