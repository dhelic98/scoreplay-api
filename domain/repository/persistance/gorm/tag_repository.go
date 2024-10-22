package gorm

import (
	"context"

	"github.com/dhelic98/scoreplay-api/domain/entity"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TagRepository struct {
	DB *gorm.DB
}

func NewPostgresSQLTagRepository(db *gorm.DB) *TagRepository {
	return &TagRepository{DB: db}
}

// Write
func (repository *TagRepository) CreateTag(ctx context.Context, tag *entity.Tag) error {
	return repository.DB.Create(tag).Error
}

// Read
func (repository *TagRepository) GetAllTags(ctx context.Context) ([]*entity.Tag, error) {
	var tags []*entity.Tag
	if err := repository.DB.Find(&tags).Error; err != nil {
		return nil, err
	}
	return tags, nil
}

func (repository *TagRepository) GetTagById(ctx context.Context, id uuid.UUID) (*entity.Tag, error) {
	var tag entity.Tag
	if err := repository.DB.First(&tag, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &tag, nil
}
