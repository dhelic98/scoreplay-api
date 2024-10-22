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
func (r *ImageRepository) CreateImage(ctx context.Context, image *entity.Image) error {
	return r.DB.Create(image).Error
}

// Read
func (r *ImageRepository) GetImageByID(ctx context.Context, id uuid.UUID) (*entity.Image, error) {
	var image entity.Image
	if err := r.DB.Preload("Tags").First(&image, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &image, nil
}

func (r *ImageRepository) GetAllImages(ctx context.Context) ([]*entity.Image, error) {
	var images []*entity.Image
	if err := r.DB.Preload("Tags").Find(&images).Error; err != nil {
		return nil, err
	}
	return images, nil
}

func (r *ImageRepository) SearchByTagName(ctx context.Context, tagName string) ([]*entity.Image, error) {
	var images []*entity.Image
	if err := r.DB.Joins("JOIN image_tags ON images.id = image_tags.image_id").
		Joins("JOIN tags ON image_tags.tag_id = tags.id").
		Where("tags.name = ?", tagName).
		Preload("Tags").
		Find(&images).Error; err != nil {
		return nil, err
	}

	return images, nil
}
