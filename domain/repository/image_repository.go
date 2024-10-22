package repository

import (
	"context"

	"github.com/dhelic98/scoreplay-api/domain/entity"
	"github.com/google/uuid"
)

type MediaRepository interface {
	//Read
	GetAllImages(ctx context.Context) ([]*entity.Image, error)
	GetImageByID(ctx context.Context, id uuid.UUID) (*entity.Image, error)
	SearchByTagName(ctx context.Context, tagName string) ([]*entity.Image, error)

	//Write
	CreateImage(ctx context.Context, image *entity.Image) error
}
