package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	Port               int
	Env                string
	Host               string
	DeployedAt         time.Time
	DBConnectionString string
}

func NewConfig() Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		log.Fatal("PORT not entered in INT format")
	}

	cfg := Config{
		Port:               port,
		Env:                os.Getenv("ENV"),
		DeployedAt:         time.Now(),
		Host:               os.Getenv("HOST"),
		DBConnectionString: os.Getenv("DB_CONNECTION_STRING"),
	}

	fmt.Println("CONFIG LOADED")

	return cfg
}
