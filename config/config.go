package config

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
	Port       string
	JwtSecret  string
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
		Port:       ":3000",
		JwtSecret:  "this_is_secret",
	}, nil
}
