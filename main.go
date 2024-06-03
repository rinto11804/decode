package main

import (
	"log"
)

func main() {
	config, err := LoadConfig()

	if err != nil {
		log.Fatal(err)
	}

	db, err := NewMongoDBStorage(config.MongoDBURI)
	if err != nil {
		log.Fatal(err)
	}

	server := NewAPIServer(":3000", db)
	server.Run()
}
