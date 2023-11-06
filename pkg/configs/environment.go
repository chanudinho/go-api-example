package configs

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

// DB_HOST=localhost
// DB_NAME=goapi
// DB_USER=user
// DB_PASSWORD=pass
// DB_PORT=3306

type EnvConfigs struct {
	DbHost     string
	DbName     string
	DbUser     string
	DbPassword string
	DbPort     int
	ApiPort    int
}

func Load() EnvConfigs {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	apiPort, err := strconv.Atoi(os.Getenv("API_PORT"))
	if err != nil {
		log.Fatal(err)
	}

	dbPort, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		log.Fatal(err)
	}

	return EnvConfigs{
		DbHost:     os.Getenv("DB_HOST"),
		DbName:     os.Getenv("DB_NAME"),
		DbUser:     os.Getenv("DB_USER"),
		DbPassword: os.Getenv("DB_PASSWORD"),
		DbPort:     dbPort,
		ApiPort:    apiPort,
	}
}
