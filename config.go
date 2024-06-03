package main

import (
	"errors"
	"os"

	"github.com/joho/godotenv"
)

var (
	ErrMongoDBURINotFound = errors.New("mongodb uri not found")
)

type Config struct {
	MongoDBURI string
}

func LoadConfig() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		return nil, err
	}
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		return nil, ErrMongoDBURINotFound
	}
	return &Config{
		MongoDBURI: uri,
	}, nil
}
