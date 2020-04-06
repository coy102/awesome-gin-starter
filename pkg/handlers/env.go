package handler

import (
    "github.com/joho/godotenv"
    "log"
)

// InitEnv environtment
func InitEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error Loading .env file")
	}
}
