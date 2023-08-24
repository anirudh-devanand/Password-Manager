package models

import (
	"log"
	"context"
	"net/http"
	"encoding/json"

	"github.com/anirudh-devanand/PwdMngr-Go/src/initialize"

	"go.mongodb.org/mongo-driver/bson"
	"github.com/gorilla/mux"
)


type Entry struct {
	Website  string `json:"Website"`
	Password string `json:"Password"`
}


// create adds an entry to the map with the given website and password
func Create(w http.ResponseWriter, r *http.Request) {

	var entry Entry

	err := json.NewDecoder(r.Body).Decode(&entry)
	if err != nil {
        http.Error(w, "Failed to parse request body", http.StatusBadRequest)
		log.Println("Failed to parse POST request body")
        return
    }

	entryDB := initialize.GetDB()

	log.Println("Adding entry...")

	if entry.Website == "" || entry.Password == "" {
		http.Error(w,"Invalid credentials entered", http.StatusBadRequest)
		log.Println("Invalid credentials entered")
		return
	} else {
		_, err := entryDB.InsertOne(context.TODO(), bson.M{"Website":entry.Website, "Password": entry.Password})
		if err!=nil{
			http.Error(w,"Error adding entry to database", http.StatusInternalServerError)
			log.Fatal("Entry not added: ",err)
			return
		}

		log.Println("Entry added")
	}
}


// read retrieves and prints either the password or website depending on the entry
func Read(w http.ResponseWriter, r *http.Request) {

	HTTPparameters := mux.Vars(r)

	var entry Entry
	entry.Website = HTTPparameters["website"]

	entryDB := initialize.GetDB()

	log.Println("Lookup requested...")

	var filter bson.M
	var result Entry

	if entry.Website == "" { 
			http.Error(w,"Invalid credentials entered", http.StatusBadRequest)
			log.Println("Invalid credentials entered")
			return
		}

	filter = bson.M{"Website" : entry.Website}

	err := entryDB.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		http.Error(w,"Error searching for entry in database", http.StatusInternalServerError)
		log.Fatal("Lookup failed: ",err)
		return
	}

	log.Println("Lookup successful")

	jsonResponse, err := json.Marshal(result)
	if err != nil {
		http.Error(w, "Error creating JSON response", http.StatusInternalServerError)
		log.Println("Error creating JSON response: ", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	_, err = w.Write(jsonResponse)
	if err != nil {
		http.Error(w, "Error writing JSON response", http.StatusInternalServerError)
		log.Println("Error creating JSON response: ",err)
		return
	}

}


// update updates the password for a given website, and if the updated password is empty, calls del()
func Update(w http.ResponseWriter, r *http.Request) {

	var entry Entry

	err := json.NewDecoder(r.Body).Decode(&entry)
	if err != nil {
        http.Error(w, "Failed to parse request body", http.StatusBadRequest)
		log.Println("Failed to parse POST request body")
        return
    }

	entryDB := initialize.GetDB()
	
	log.Println("Updating entry...")

	if entry.Website == ""{
		http.Error(w,"Invalid credentials entered", http.StatusBadRequest)
		log.Println("Invalid credentials entered")
		return
	}

	if entry.Password == "" {
		Delete(w, r)
	}

	filter := bson.M{"Website" : entry.Website}
	update := bson.M{"$set": bson.M{"Password" : entry.Password}}

	_, err = entryDB.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		http.Error(w,"Error updating entry in database", http.StatusInternalServerError)
		log.Fatal("Update failed: ",err)
		return
	}

	log.Println("Update successful")
}


// delete deletes an entry from the map for the given website
func Delete(w http.ResponseWriter, r *http.Request) {
	HTTPparameters := mux.Vars(r)
	website := HTTPparameters["website"]

	entryDB := initialize.GetDB()

	log.Println("Deleting entry...")

	filter := bson.M{"Website": website}

	_, err := entryDB.DeleteOne(context.TODO(), filter)
	if err != nil {
		http.Error(w,"Error deleting entry in database", http.StatusInternalServerError)
		log.Println("Deletion unsuccessful: ",err)
		return
	}

	log.Println("Deletion successful")
}
