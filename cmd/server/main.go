package main

import (
	"fmt"
	"net/http"

	"github.com/dhelic98/scoreplay-api/application/router"
	"github.com/dhelic98/scoreplay-api/cmd/config"
	_ "github.com/dhelic98/scoreplay-api/docs"
	customhttp "github.com/dhelic98/scoreplay-api/interface/http"
	httpSwagger "github.com/swaggo/http-swagger"
)

//	@title			Scoreplay Media API
//	@version		1.0
//	@description	Simple REST API for CRUD operations on Media and Tag entities.

//	@host		localhost:4200
//	@BasePath	/v1

func main() {
	config := config.GetConfigInstance()
	db := InitiateDatabaseConnection()
	router := router.SetupRoutes(db)

	router.Handle("/swagger/", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:4200/swagger/doc.json"),
	))

	server := http.Server{
		Addr:    fmt.Sprintf(":%v", config.Port),
		Handler: customhttp.Logger(router),
	}

	fmt.Printf("Starting server on PORT:%v\n", config.Port)
	server.ListenAndServe()

}
