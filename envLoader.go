package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	GROUP_IRON_BASE_URL string
	GROUP_IRON_TOKEN    string
	PORT                string 
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Error loading .env file, proceeding with environment variables if present.")
		return nil, err
	}

	log.Println("Loading .env")
	return &Config{
		os.Getenv("GROUP_IRON_BASE_URL"),
		os.Getenv("GROUP_IRON_TOKEN"),
		":" + os.Getenv("HTTP_PORT"),
	}, nil
}
