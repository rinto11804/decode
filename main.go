package main

import (
	"decode/config"
	"log"
)

func main() {
	config, err := config.LoadConfig()

	if err != nil {
		log.Fatal(err)
	}
	client, err := NewMongoDBStorage(config.MongoDBURI)
	if err != nil {
		log.Fatal(err)
	}

	server := NewAPIServer(config, client)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
