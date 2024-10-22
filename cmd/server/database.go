package main

import (
	"fmt"
	"log"

	"github.com/dhelic98/scoreplay-api/cmd/config"
	"github.com/dhelic98/scoreplay-api/domain/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitiateDatabaseConnection() *gorm.DB {
	fmt.Println("OPENING DATABASE CONNECTION")
	db, err := gorm.Open(postgres.Open(config.GetConfigInstance().DBConnectionString), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}

	fmt.Println("DATABASE MIGRATION")
	// Auto-migrate tables
	db.AutoMigrate(&entity.Tag{}, &entity.Image{})
	fmt.Println("CONNECTION SUCCESSFULLY OPENED")
	return db
}
