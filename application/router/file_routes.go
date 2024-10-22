package router

import (
	"net/http"

	"github.com/dhelic98/scoreplay-api/application/handler"
)

func registerFileRoutes(router *http.ServeMux, fileHandler *handler.FileHandler) {
	router.HandleFunc("GET /file/{fileID}", fileHandler.ServeImageFile)
}
