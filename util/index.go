package util

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetEnvVariable (key string) string {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error Loading ENV Variables!!!")
	}

	return os.Getenv(key)
}

func GetEnvConfig (keys ...string) map[string]string {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error Loading ENV Variables!!!")
	}

	var configMap = make(map[string]string)

	for _, key := range keys {
		variable := os.Getenv(key)

		configMap[key] = variable
	}

	return configMap
}