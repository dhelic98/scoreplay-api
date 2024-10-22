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

func (handler *FileHandler) ServeImageFile(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("fileID")

	filePath := handler.FileService.GetFilePath(idStr)

	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		http.Error(w, "File not found", http.StatusNotFound)
		return
	}

	http.ServeFile(w, r, filePath)
}
