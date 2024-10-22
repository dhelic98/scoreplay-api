package handler

import (
	"encoding/json"
	"net/http"
	"path/filepath"

	"github.com/dhelic98/scoreplay-api/application/dto"
	"github.com/dhelic98/scoreplay-api/application/service"
	"github.com/dhelic98/scoreplay-api/interface/enum"
	"github.com/google/uuid"
)

type ImageHandler struct {
	ImageService service.IImageService
	FileService  service.IFileService
	TagService   service.ITagService
}

func NewImageHandler(imageService service.IImageService, tagService service.ITagService, fileService service.IFileService) *ImageHandler {
	return &ImageHandler{FileService: fileService, ImageService: imageService, TagService: tagService}
}

func (handler *ImageHandler) CreateImageHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		http.Error(w, "Failed to parse form", http.StatusBadRequest)
		return
	}

	name := r.FormValue("name")
	if len(name) < 1 {
		http.Error(w, "Media Name is required", http.StatusBadRequest)
		return
	}

	file, fileHeader, err := r.FormFile("image")
	if err != nil || fileHeader == nil {
		http.Error(w, "Image file is required", http.StatusBadRequest)
		return
	}
	defer file.Close()

	fileExt := filepath.Ext(fileHeader.Filename)
	if enum.IsAllowed(fileExt) == false {
		http.Error(w, "File format unsupported", http.StatusBadRequest)
		return
	}

	tagsJSONString := r.FormValue("tags")
	if tagsJSONString == "" {
		http.Error(w, "Tags are required", http.StatusBadRequest)
		return
	}

	tagIDs, err := handler.TagService.ParseMultipartFormToUUID(tagsJSONString)
	if err != nil {
		http.Error(w, "Failed to parse tags", http.StatusInternalServerError)
		return
	}

	fileUrl, err := handler.FileService.UploadFile(&file, fileHeader)
	if err != nil {
		http.Error(w, "Failed to save image to server", http.StatusInternalServerError)
		return
	}

	err = handler.ImageService.CreateImage(r.Context(), &dto.CreateImageDTO{
		Name: name,
		URL:  fileUrl,
		Tags: tagIDs,
	})
	if err != nil {
		http.Error(w, "Failed to create image", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (handler *ImageHandler) GetAllImagesHandler(w http.ResponseWriter, r *http.Request) {
	images, err := handler.ImageService.GetAllImages(r.Context())
	if err != nil {
		http.Error(w, "Images not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(images)
}

func (handler *ImageHandler) GetImageByIDHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, "Invalid UUID", http.StatusBadRequest)
		return
	}

	image, err := handler.ImageService.GetImageById(r.Context(), id)
	if err != nil {
		http.Error(w, "Image with ID provided not found ", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(image)
}

func (handler *ImageHandler) SearchByTag(w http.ResponseWriter, r *http.Request) {
	tagStr := r.PathValue("tagName")

	if len(tagStr) == 0 {
		http.Error(w, "No tag provided", http.StatusBadRequest)
	}

	images, err := handler.ImageService.SearchImagesByTagName(r.Context(), tagStr)
	if err != nil {
		http.Error(w, "Images not found with tag name provided", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(images)
}
