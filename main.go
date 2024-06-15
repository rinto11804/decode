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
	db, err := NewMongoDBStorage(config.MongoDBURI)
	if err != nil {
		log.Fatal(err)
	}

	server := NewAPIServer(config, db)
	server.Run()
}
