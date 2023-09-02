package initialize

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)


// Define global variables to serve getter functions
var (
	mongoClient *mongo.Client
	mongoDB     *mongo.Collection
)

// initClient initializes the MongoDB client and establishes a connection to the database.
func initClient() {
	// Get the MongoDB host from the environment variable passed during main.go runtime
	host := os.Getenv("HOST")

	// Check if the environment variable is set
	if host == "" {
		log.Fatal("HOST is not set")
	} else {
		log.Printf("HOST is set to %s\n", host)
	}


	// Set MongoDB's Unique Resource Identifier dynamically
	uri := "mongodb://" + host + ":27017"

	clientOptions := options.Client().ApplyURI(uri)
	var err error

	// Connect to MongoDB
	mongoClient, err = mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal("Error connecting to MongoDB: ", err)
	}

	// Ping the MongoDB server to check if it's responsive
	err = mongoClient.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal("Error connecting to MongoDB: ", err)
	}

	// If the Ping operation was successful, the MongoDB server is responsive
	log.Println("Connection to MongoDB initialized")
}

// getClient is a getter function that returns the MongoDB client.
func getClient() *mongo.Client {
	return mongoClient
}

// InitDB initializes the MongoDB database and collection for the application.
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

// GetDB is a getter function that returns the MongoDB collection.
func GetDB() *mongo.Collection {
	return mongoDB
}
