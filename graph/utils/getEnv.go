package utils

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var env map[string]string = nil

// InitDotEnv initiates .env
func InitDotEnv() {
	err := godotenv.Load()
	if err != nil {
		handleErrors(err, "Error loading .env file")
		return
	}

	loadedEnv, err := godotenv.Read()
	if err != nil {
		handleErrors(err, "Error reading .env file")
		return
	}

	env = loadedEnv
}

// GetEnv gets the environment variable from either .env or os
func GetEnv(key string) string {
	if nil != env {
		return env[key]
	}
	return os.Getenv(key)
}

func handleErrors(err error, message string) {
	if os.Getenv("APP_ENV") != "production" {
		log.Fatal(message)
	}
}
