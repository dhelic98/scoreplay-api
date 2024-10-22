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
func (r *TagRepository) CreateTag(ctx context.Context, tag *entity.Tag) error {
	return r.DB.Create(tag).Error
}

// Read
func (r *TagRepository) GetAllTags(ctx context.Context) ([]*entity.Tag, error) {
	var tags []*entity.Tag
	if err := r.DB.Find(&tags).Error; err != nil {
		return nil, err
	}
	return tags, nil
}

func (r *TagRepository) GetTagById(ctx context.Context, id uuid.UUID) (*entity.Tag, error) {
	var tag entity.Tag
	if err := r.DB.First(&tag, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &tag, nil
}

func (r *TagRepository) GetByIDs(ctx context.Context, ids []uuid.UUID) ([]entity.Tag, error) {
	var tags []entity.Tag
	if err := r.DB.Where("id IN ?", ids).Find(&tags).Error; err != nil {
		return nil, err
	}
	return tags, nil
}
