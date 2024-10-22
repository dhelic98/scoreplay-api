package config

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
	CurrentAPIVersion  string
	FileHostUrl        string
}

var singletonConfig *Config

func GetConfigInstance() *Config {
	if singletonConfig == nil {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}

		port, err := strconv.Atoi(os.Getenv("PORT"))
		if err != nil {
			log.Fatal("PORT not entered in INT format")
		}

		singletonConfig = &Config{
			Port:               port,
			Env:                os.Getenv("ENV"),
			DeployedAt:         time.Now(),
			Host:               os.Getenv("HOST"),
			DBConnectionString: os.Getenv("DB_CONNECTION_STRING"),
			CurrentAPIVersion:  os.Getenv("CURRENT_API_VERSION"),
			FileHostUrl:        os.Getenv("FILE_HOST_URL"),
		}

		fmt.Println("CONFIG LOADED")
	}

	return singletonConfig
}
