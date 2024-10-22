package router

import (
	"net/http"

	"github.com/dhelic98/scoreplay-api/application/handler"
	"github.com/dhelic98/scoreplay-api/application/service"
	persistance "github.com/dhelic98/scoreplay-api/infrastructure/persistance/gorm"
	"gorm.io/gorm"
)

func SetupRoutes(db *gorm.DB) *http.ServeMux {

	router := http.NewServeMux()
	//Create repositories
	imageRepository := &persistance.ImageRepository{DB: db}
	tagRepository := &persistance.TagRepository{DB: db}

	//Create Services
	tagService := &service.TagService{Respository: tagRepository}
	imageService := &service.MediaService{Respository: imageRepository, TagRepository: tagRepository}

	//Create handlers
	tagHandler := &handler.TagHandler{Service: tagService}
	imageHandler := &handler.MediaHandler{Service: imageService}

	//Registering routes
	registerTagRoutes(router, tagHandler)
	registerImageRoutes(router, imageHandler)

	//Adding v1 prefix for versioning
	v1 := http.NewServeMux()
	v1.Handle("/v1/", http.StripPrefix("/v1", router))

	return v1
}

func registerTagRoutes(router *http.ServeMux, tagHandler *handler.TagHandler) {

	router.HandleFunc("GET /tags", tagHandler.GetAllTagsHandler)
	router.HandleFunc("GET /tags/{id}", tagHandler.GetTagByIDHandler)
	router.HandleFunc("POST /tags", tagHandler.CreateTagHandler)
}

func registerImageRoutes(router *http.ServeMux, mediaHandler *handler.MediaHandler) {
	router.HandleFunc("GET /media", mediaHandler.GetAllImagesHandler)
	router.HandleFunc("GET /media/{id}", mediaHandler.GetImageByIDHandler)
	router.HandleFunc("POST /media", mediaHandler.CreateImageHandler)
	router.HandleFunc("GET /media/filter/{tagName}", mediaHandler.SearchByTag)
	router.HandleFunc("GET /media/file/{fileID}", mediaHandler.ServeImageFile)
}
