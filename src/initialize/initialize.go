package initialize

import (
	"context"
	"log"
	"os"
	
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	mongoClient *mongo.Client
	mongoDB     *mongo.Collection
)

func initClient() {

	host := os.Getenv("LOCALHOST")

    // Check if the environment variable is set
    if host == "" {
        log.Fatal("HOST is not set")
    } else {
        log.Printf("HOST is set to %s\n", host)
    }

	uri := "mongodb://" + host + ":27017"

	clientOptions := options.Client().ApplyURI(uri)
	var err error

	// Connect to MongoDB
	mongoClient, err = mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal("Error connecting to mongoDB: ", err)
	}

	// Ping the MongoDB server to check if it's responsive
	err = mongoClient.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal("Error connecting to mongoDB: ", err)
	}

	// If the Ping operation was successful, the MongoDB server is responsive
	log.Println("Connection to MongoDB inititalized")

}

func getClient() *mongo.Client {
	return mongoClient
}

func InitDB() {

	initClient()
	DBclient := getClient()
	if DBclient != nil {
		mongoDB = DBclient.Database("PswdMngr").Collection("EntryDB")

		log.Println("Created EntryDB database for PswdMngr")
	} else {

		log.Fatal("Error creating EntryDB database for PswdMngr")
	}

}

func GetDB() *mongo.Collection {
	return mongoDB
}
