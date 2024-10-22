package router

import (
	"net/http"

	"github.com/dhelic98/scoreplay-api/application/handler"
	"github.com/dhelic98/scoreplay-api/application/service"
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

	//Adding v1 prefix for versioning
	v1 := http.NewServeMux()
	v1.Handle("/v1/", http.StripPrefix("/v1", router))

	return v1
}
