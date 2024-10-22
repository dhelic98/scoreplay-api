package dto

import (
	"github.com/dhelic98/scoreplay-api/domain/entity"
	"github.com/google/uuid"
)

type CreateTagDTO struct {
	Name string `json:"name" validate:"required"`
}

type GetTagDTO struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

func ToTagDTO(tag entity.Tag) GetTagDTO {
	return GetTagDTO{
		ID:   tag.ID,
		Name: tag.Name,
	}
}
