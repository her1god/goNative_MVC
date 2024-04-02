package config

import (
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	DB_USER     string
	DB_PASSWORD string
	DB_NAME     string
	DB_HOST     string
	DB_PORT     int
)

func LoadEnv() error {
	err := godotenv.Load()
	if err != nil {
		return err
	}

	DB_USER = os.Getenv("DB_USER")
	DB_PASSWORD = os.Getenv("DB_PASSWORD")
	DB_NAME = os.Getenv("DB_NAME")
	DB_HOST = os.Getenv("DB_HOST")
	DB_PORT, _ = strconv.Atoi(os.Getenv("DB_PORT"))

	return nil
}
