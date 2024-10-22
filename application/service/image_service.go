package service

import (
	"context"
	"fmt"

	"github.com/dhelic98/scoreplay-api/application/dto"
	"github.com/dhelic98/scoreplay-api/domain/entity"
	"github.com/dhelic98/scoreplay-api/domain/repository"
	"github.com/google/uuid"
)

type ImageService struct {
	Respository repository.ImageRepository
}

func NewImageService(imageRepository repository.ImageRepository) *ImageService {
	return &ImageService{Respository: imageRepository}
}

func (imageService *ImageService) CreateImage(ctx context.Context, imageDTO *dto.CreateImageDTO) error {
	tags := make([]entity.Tag, len(imageDTO.Tags))
	for i, tagUUID := range imageDTO.Tags {
		tags[i] = entity.Tag{
			ID: tagUUID,
		}
	}

	image := entity.Image{
		ID:   uuid.New(),
		Name: imageDTO.Name,
		Tags: tags,
		URL:  imageDTO.URL,
	}

	return imageService.Respository.CreateImage(ctx, &image)
}

func (imageService *ImageService) GetAllImages(ctx context.Context) ([]*dto.GetImageDTO, error) {
	images, err := imageService.Respository.GetAllImages(ctx)
	if err != nil {
		fmt.Println("error in here ")
		return nil, err
	}

	imagesDTO := make([]*dto.GetImageDTO, len(images))

	for i, image := range images {
		tagNames := make([]string, len(image.Tags))
		for i, tag := range image.Tags {
			tagNames[i] = tag.Name
		}
		imagesDTO[i] = dto.ToGetImageDTO(image)
	}

	return imagesDTO, nil
}

func (imageService *ImageService) GetImageById(ctx context.Context, id uuid.UUID) (*dto.GetImageDTO, error) {
	image, err := imageService.Respository.GetImageByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return dto.ToGetImageDTO(image), nil
}

func (imageService *ImageService) SearchImagesByTagName(ctx context.Context, tagName string) ([]*dto.GetImageDTO, error) {
	images, err := imageService.Respository.SearchByTagName(ctx, tagName)
	if err != nil {
		return nil, err
	}

	imageDTOs := make([]*dto.GetImageDTO, len(images))
	for i, image := range images {
		imageDTOs[i] = dto.ToGetImageDTO(image)
	}
	return imageDTOs, nil
}
