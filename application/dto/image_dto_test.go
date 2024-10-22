package dto_test

import (
	"testing"

	"github.com/dhelic98/scoreplay-api/application/dto"
	"github.com/dhelic98/scoreplay-api/domain/entity"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestToGetImageDTO(t *testing.T) {
	id := uuid.New()
	tagId1 := uuid.New()
	tagId2 := uuid.New()

	image := &entity.Image{
		ID:   id,
		Name: "Test Image",
		URL:  "/uploads/image.png",
		Tags: []entity.Tag{
			{ID: tagId1, Name: "Tag1"},
			{ID: tagId2, Name: "Tag2"},
		},
	}

	imageDTO := dto.ToGetImageDTO(image)

	assert.Equal(t, image.ID, imageDTO.ID)
	assert.Equal(t, image.Name, imageDTO.Name)
	assert.Equal(t, image.URL, imageDTO.URL)
	assert.Equal(t, []string{"Tag1", "Tag2"}, imageDTO.Tags)
}
