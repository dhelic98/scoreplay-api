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

// CreateImageHandler creates a new media/image.
//
//	@Summary		Create a new image
//	@Description	Create a new image with the specified name tags
//	@Tags			images
//	@Accept			mpfd
//
//	@Param			name	formData	string		true	"Image Name"
//	@Param			tags	formData	[]string	true	"Array of tag IDs (e.g., tags=['id1','id2'])"
//	@Param			image	formData	file		true	"The image file to upload"
//
//	@Success		201
//	@Failure		400
//	@Failure		500
//	@Router			/media [post]
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

// GetAllImagesHandler get all images.
//
//	@Summary	Get all image entities
//	@Tags		images
//	@Produce	json
//	@Success	200
//	@Failure	500
//	@Router		/media [get]
func (handler *ImageHandler) GetAllImagesHandler(w http.ResponseWriter, r *http.Request) {
	images, err := handler.ImageService.GetAllImages(r.Context())
	if err != nil {
		http.Error(w, "Images not found", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(images)
}

// GetImageByIDHandler get images by UUID.
//
//	@Summary	Get images by ID
//	@Tags		images
//	@Produce	json
//
//	@Param		id	path	string	true	"Image UUID"
//
//	@Success	200
//	@Failure	400
//	@Failure	404
//	@Router		/media/{id} [get]
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

// SearchByTagHandler get images by Tag name.
//
//	@Summary	Get images by tag name
//	@Tags		images
//	@Produce	json
//
//	@Param		tagName	path	string	true	"Tag name"
//
//	@Success	200
//	@Failure	400
//	@Failure	404
//	@Router		/media/filter/{tagName} [get]
func (handler *ImageHandler) SearchByTagHandler(w http.ResponseWriter, r *http.Request) {
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
