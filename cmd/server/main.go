package main

import (
	"fmt"
	"net/http"

	"github.com/dhelic98/scoreplay-api/application/router"
	customhttp "github.com/dhelic98/scoreplay-api/interface/http"
)

func main() {
	config := NewConfig()
	db := InitiateDatabaseConnection(config.DBConnectionString)
	router := router.SetupRoutes(db)

	server := http.Server{
		Addr:    fmt.Sprintf(":%v", config.Port),
		Handler: customhttp.Logger(router),
	}

	fmt.Printf("Starting server on PORT:%v\n", config.Port)
	server.ListenAndServe()

}
