package dto

import (
	"github.com/dhelic98/scoreplay-api/domain/entity"
	"github.com/google/uuid"
)

type CreateImageDTO struct {
	Name string      `json:"name"`
	Tags []uuid.UUID `json:"tags"`
	URL  string      `json:"url"`
}

type GetImageDTO struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
	Tags []string  `json:"tags"`
	URL  string    `json:"url"`
}

func ToGetImageDTO(image *entity.Image) *GetImageDTO {
	var tags []string = make([]string, len(image.Tags))
	for i, tag := range image.Tags {
		tags[i] = tag.Name
	}
	return &GetImageDTO{
		ID:   image.ID,
		Name: image.Name,
		URL:  image.URL,
		Tags: tags,
	}
}
