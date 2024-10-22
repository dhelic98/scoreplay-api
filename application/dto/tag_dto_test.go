package dto_test

import (
	"testing"

	"github.com/dhelic98/scoreplay-api/application/dto"
	"github.com/dhelic98/scoreplay-api/domain/entity"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestToGetTagDTO(t *testing.T) {
	id := uuid.New()

	tag := &entity.Tag{
		ID:   id,
		Name: "Test Image",
	}

	tagDTO := dto.ToTagDTO(tag)

	assert.Equal(t, tag.ID, tagDTO.ID)
	assert.Equal(t, tag.Name, tagDTO.Name)

}
