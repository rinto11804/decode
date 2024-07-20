package main

import (
	"decode/config"
	"log"
)

//	@title			Swagger Example API
//	@version		1.0
//	@description	This is a sample server celler server.
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@host		localhost:3000
//	@BasePath	/api/v1

//	@securityDefinitions.apikey	Bearer
//	@in							header
//	@name						Authorization

//	@externalDocs.description	OpenAPI
//	@externalDocs.url			https://swagger.io/resources/open-api/
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
