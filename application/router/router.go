package router

import (
	"fmt"
	"net/http"

	"github.com/dhelic98/scoreplay-api/application/handler"
	"github.com/dhelic98/scoreplay-api/application/service"
	"github.com/dhelic98/scoreplay-api/cmd/config"
	persistance "github.com/dhelic98/scoreplay-api/domain/repository/persistance/gorm"
	"gorm.io/gorm"
)

func SetupRoutes(db *gorm.DB) *http.ServeMux {

	router := http.NewServeMux()
	//Create repositories
	imageRepository := &persistance.ImageRepository{DB: db}
	tagRepository := &persistance.TagRepository{DB: db}

	//Create Services
	tagService := &service.TagService{Respository: tagRepository}
	imageService := &service.ImageService{Respository: imageRepository}
	fileService := &service.FileService{}

	//Create handlers
	tagHandler := &handler.TagHandler{Service: tagService}
	imageHandler := &handler.ImageHandler{ImageService: imageService, FileService: fileService}
	fileHandler := &handler.FileHandler{FileService: fileService}

	//Registering routes
	registerTagRoutes(router, tagHandler)
	registerImageRoutes(router, imageHandler)
	registerFileRoutes(router, fileHandler)

	//Adding versionRouter prefix for versioning
	versionRouter := http.NewServeMux()
	versionRouter.Handle(fmt.Sprintf("/%s/", config.GetConfigInstance().CurrentAPIVersion),
		http.StripPrefix(fmt.Sprintf("/%s", config.GetConfigInstance().CurrentAPIVersion), router))

	return versionRouter
}
