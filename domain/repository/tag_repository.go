package repository

import (
	"context"

	"github.com/dhelic98/scoreplay-api/domain/entity"
	"github.com/google/uuid"
)

type TagRepository interface {
	//Read
	GetAllTags(ctx context.Context) ([]*entity.Tag, error)
	GetTagById(ctx context.Context, id uuid.UUID) (*entity.Tag, error)

	//Write
	CreateTag(ctx context.Context, tag *entity.Tag) error
}
