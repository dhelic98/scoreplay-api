package router

import (
	"net/http"

	"github.com/dhelic98/scoreplay-api/application/handler"
)

func registerImageRoutes(router *http.ServeMux, imageHandler *handler.ImageHandler) {
	router.HandleFunc("GET /media", imageHandler.GetAllImagesHandler)
	router.HandleFunc("GET /media/{id}", imageHandler.GetImageByIDHandler)
	router.HandleFunc("POST /media", imageHandler.CreateImageHandler)
	router.HandleFunc("GET /media/filter/{tagName}", imageHandler.SearchByTagHandler)

}
