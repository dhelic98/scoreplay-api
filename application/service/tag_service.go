package service

import (
	"context"

	"github.com/dhelic98/scoreplay-api/application/dto"
	"github.com/dhelic98/scoreplay-api/domain/entity"
	"github.com/dhelic98/scoreplay-api/domain/repository"
	"github.com/google/uuid"
)

type TagService struct {
	Respository repository.TagRepository
}

func NewTagService(respository repository.TagRepository) *TagService {
	return &TagService{Respository: respository}
}

func (tagService *TagService) CreateTag(ctx context.Context, createTagDto dto.CreateTagDTO) error {
	tag := entity.Tag{
		ID:   uuid.New(),
		Name: createTagDto.Name,
	}
	return tagService.Respository.CreateTag(ctx, &tag)
}

func (tagService *TagService) GetAllTags(ctx context.Context) ([]*dto.GetTagDTO, error) {
	tags, err := tagService.Respository.GetAllTags(ctx)
	if err != nil {
		return nil, err
	}
	tagDTOs := make([]*dto.GetTagDTO, len(tags))

	for i, v := range tags {
		tagDTOs[i] = &dto.GetTagDTO{
			ID:   v.ID,
			Name: v.Name,
		}
	}
	return tagDTOs, nil
}

func (tagService *TagService) GetTagById(ctx context.Context, id uuid.UUID) (*dto.GetTagDTO, error) {
	tag, err := tagService.Respository.GetTagById(ctx, id)
	if err != nil {
		return nil, err
	}
	return &dto.GetTagDTO{
		ID:   tag.ID,
		Name: tag.Name,
	}, nil
}