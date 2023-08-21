package main

import(
	"fmt"
	"github.com/anirudh-devanand/PwdMngr-Go/src/models"
	"log"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)



func entryDB() (*mongo.Client){
	// Set client options to specify the MongoDB connection URI
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Connect to MongoDB
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	defer client.Disconnect(context.Background()) // Close the connection when done

	// Ping the MongoDB server to check if it's responsive
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	// If the Ping operation was successful, the MongoDB server is responsive
	log.Println("Connected to MongoDB")

	return client


}

func main(){

	DBclient := entryDB();

	entryDB := DBclient.Database("PswdMngr").Collection("EntryDB")

	log.Println("Created database for PswdMngr")
	
	fmt.Println("SYNTAX: <OPERATION> <WEBSITE>* <PASSWORD>** \nOPERATIONS:\nCREATE\tREAD\tUPDATE\tDELETE\texit\n*Using the READ operation with website as \"NULL\" and a password will return webiste\n** Using UPDATE operation with passowrd as \"\" will delete website entry")
	for {
		var (
			operation string
			entry models.Entry
			website string
			password string
		)

		fmt.Scanf("%s %s %s", &operation, &website, &password)

		if operation == "exit"{
			err := DBclient.Disconnect(context.TODO())

			if err != nil {
    			log.Fatal(err)
			}

			fmt.Println("Connection to MongoDB closed")
			break;
		}

		entry.Website=website
		entry.Password=password
		entry.EntryOperations(operation, entryDB)

	}

}