package handler

import (
	"net/http"
	"os"

	"github.com/dhelic98/scoreplay-api/application/service"
)

type FileHandler struct {
	FileService service.IFileService
}

func NewFileHandler(fileService service.IFileService) *FileHandler {
	return &FileHandler{FileService: fileService}
}

// ServeImageFile serves image file
//
//	@Summary	Get image file by image UUID
//	@Produce	png
//	@Produce	jpeg
//
//	@Param		fileID	path	string	true	"Image file UUID"
//
//	@Tags		files
//	@Success	200
//	@Failure	404
func (handler *FileHandler) ServeImageFile(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("fileID")

	filePath := handler.FileService.GetFilePath(idStr)

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		http.Error(w, "File not found", http.StatusNotFound)
		return
	}

	http.ServeFile(w, r, filePath)
}
