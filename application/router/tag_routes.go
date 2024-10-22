package router

import (
	"net/http"

	"github.com/dhelic98/scoreplay-api/application/handler"
)

func registerTagRoutes(router *http.ServeMux, tagHandler *handler.TagHandler) {
	router.HandleFunc("GET /tags", tagHandler.GetAllTagsHandler)
	router.HandleFunc("GET /tags/{id}", tagHandler.GetTagByIDHandler)
	router.HandleFunc("POST /tags", tagHandler.CreateTagHandler)
}
