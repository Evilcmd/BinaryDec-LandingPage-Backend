package main

import (
	"log"
	"os"

	"github.com/Evilcmd/Hackup-backend/internal/apis"
	mongodb "github.com/Evilcmd/Hackup-backend/internal/mongoDB"
	"github.com/Evilcmd/Hackup-backend/internal/server"
	"github.com/joho/godotenv"
)

func main() {

	godotenv.Load()

	mongoUrl := os.Getenv("MONGO_DB_URL")
	if mongoUrl == "" {
		log.Fatal("error getting mongodb connection url")
	}
	mongoClient, err := mongodb.NewMongoDbClient(mongoUrl)
	if err != nil {
		log.Fatal("error creating mongodb client: ", err.Error())
	}

	apiCfg := &apis.ApiConfig{
		UserDbClient: mongoClient,
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "80"
	}
	server := server.NewServer(port, apiCfg)

	log.Println("Starting Server")
	if err := server.ListenAndServe(); err != nil {
		log.Fatal("error starting the server")
	}
}
