package main

import (
	"github.com/joho/godotenv"
	"os"
)

const (
	envFile = ".env"
)

// TODO: Need to Refactoring to factory pattern for multi config format (e.g .yaml file or anything)
func CreateDBConfigFromEnv() *DBConfig {
	_ = godotenv.Load(envFile)
	return &DBConfig{
		Host: os.Getenv("MONGO_DB_HOST"),
		Username: os.Getenv("MONGO_DB_USERNAME"),
		Password: os.Getenv("MONGO_DB_PASSWORD"),
		Port: os.Getenv("MONGO_DB_PORT"),
	}
}
