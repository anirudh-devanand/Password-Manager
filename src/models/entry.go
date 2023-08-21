package models

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// Entry represents a website-password entry
type Entry struct {
	Website  string
	Password string
}


// EntryOps performs various operations on an Entry based on the given operation
func (entry Entry) EntryOperations(operation string, entryDB *mongo.Collection) {
	
	switch operation {
	case "CREATE":
		create(entry, entryDB)
	case "READ":
		read(entry, entryDB)
	case "UPDATE":
		update(entry, entryDB)
	case "DELETE":
		del(entry.Website, entryDB)
	}
}

// create adds an entry to the map with the given website and password
func create(entry Entry, entryDB *mongo.Collection) {

	log.Println("Adding entry...")
	if entry.Website == "" || entry.Password == "" {
		fmt.Printf("ERROR: Enter valid credentials\n")
		return
	} else {


		newEntry := bson.M{entry.Website : entry.Password}
		_, err := entryDB.InsertOne(context.TODO(), newEntry)

		if err!=nil{
			log.Fatal("Entry not added: ",err)

			return
		}

		log.Println("Entry added")
	}
}

// read retrieves and prints either the password or website depending on the entry
func read(entry Entry, entryDB *mongo.Collection) {

	log.Println("Lookup requested...")
	var filter bson.M
	var result Entry

	if entry.Password == "" {

		if entry.Website == "" { 
			fmt.Printf("ERROR: Enter valid credentials\n")
			return
		}

		filter = bson.M{entry.Website: bson.M{"$exists":true}}
	} else {
		filter = bson.M{"" : entry.Password}
	}

	err := entryDB.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		log.Fatal("Lookup failed: ",err)
		return
	}

	log.Println("Lookup successful")
	fmt.Printf("Found entry: %+v\n", result)
}

// update updates the password for a given website, and if the updated password is empty, calls del()
func update(entry Entry, entryDB *mongo.Collection) {

	log.Println("Updating entry...")

	if entry.Website == ""{
		fmt.Printf("ERROR: Enter valid credentials\n")
		return
	}

	if entry.Password == "" {
		del(entry.Website, entryDB)
	}

	filter := bson.M{entry.Website : bson.M{"$exists":true}}
	update := bson.M{"$set": bson.M{entry.Website : entry.Password}}

	_, err := entryDB.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal("Update unsuccessful: ",err)
		return
	}

	log.Println("Update successful")
}

// del deletes an entry from the map for the given website
func del(website string, entryDB *mongo.Collection) {

	log.Println("Deleting entry...")

	filter := bson.M{website : bson.M{"$exists":true}}

	_, err := entryDB.DeleteOne(context.TODO(), filter)
	if err != nil {
		log.Println("Deletion unsuccessful: ",err)
		return
	}

	log.Println("Deletion successful")
}

