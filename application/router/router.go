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
	imageRepository := persistance.NewPostgresSQLImageRepository(db)
	tagRepository := persistance.NewPostgresSQLTagRepository(db)

	//Create Services
	tagService := service.NewTagService(tagRepository)
	imageService := service.NewImageService(imageRepository)
	fileService := service.NewFileService()

	//Create handlers
	tagHandler := handler.NewTagHandler(tagService)
	imageHandler := handler.NewImageHandler(imageService, tagService, fileService)
	fileHandler := handler.NewFileHandler(fileService)

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
