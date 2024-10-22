package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/dhelic98/scoreplay-api/application/dto"
	"github.com/dhelic98/scoreplay-api/application/service"
	"github.com/google/uuid"
)

type MediaHandler struct {
	Service *service.MediaService
}

func (handler *MediaHandler) CreateImageHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		http.Error(w, "Failed to parse form", http.StatusBadRequest)
		return
	}

	name := r.FormValue("name")
	file, fileHeader, err := r.FormFile("image")
	if err != nil || fileHeader == nil {
		http.Error(w, "Image file is required", http.StatusBadRequest)
		return
	}
	defer file.Close()

	var tagIDs []uuid.UUID
	tagIDStrs := r.Form["tags"]
	for _, tagIDStr := range tagIDStrs {
		tagID, err := uuid.Parse(tagIDStr)
		if err != nil {
			http.Error(w, "Invalid tag ID", http.StatusBadRequest)
			return
		}
		tagIDs = append(tagIDs, tagID)
	}

	fileID := uuid.New().String()
	filePath := "./uploads/" + fileID + ".jpg"
	fileBytes, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "Failed to read file data", http.StatusInternalServerError)
		return
	}
	err = os.WriteFile(filePath, fileBytes, os.ModePerm)
	if err != nil {
		http.Error(w, "Failed to save image file", http.StatusInternalServerError)
		return
	}

	FILE_HOST_URL := os.Getenv("FILE_HOST_URL")

	fileUrl := fmt.Sprintf("%s/images/file/%s", FILE_HOST_URL, fileID)
	createImageDTO := dto.CreateImageDTO{
		Name: name,
		URL:  fileUrl,
		Tags: tagIDs,
	}
	err = handler.Service.CreateImage(r.Context(), createImageDTO)
	if err != nil {
		http.Error(w, "Failed to create image", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (handler *MediaHandler) GetAllImagesHandler(w http.ResponseWriter, r *http.Request) {
	images, err := handler.Service.GetAllImages(r.Context())
	if err != nil {
		http.Error(w, "Images not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(images)
}

func (handler *MediaHandler) GetImageByIDHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, "Invalid UUID", http.StatusBadRequest)
		return
	}

	image, err := handler.Service.GetImageById(r.Context(), id)
	if err != nil {
		http.Error(w, "Image with ID provided not found ", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(image)
}

func (handler *MediaHandler) SearchByTag(w http.ResponseWriter, r *http.Request) {
	tagStr := r.PathValue("tagName")

	if len(tagStr) == 0 {
		http.Error(w, "No tag provided", http.StatusBadRequest)
	}

	images, err := handler.Service.SearchImagesByTagName(r.Context(), tagStr)
	if err != nil {
		http.Error(w, "Images not found with tag name provided", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(images)
}

func (h *MediaHandler) ServeImageFile(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("fileID")

	filePath := "./uploads/" + idStr + ".jpg"

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		http.Error(w, "File not found", http.StatusNotFound)
		return
	}

	http.ServeFile(w, r, filePath)
}
