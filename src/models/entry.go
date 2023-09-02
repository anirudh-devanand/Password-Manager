package models

import (
	"log"
	"context"
	"net/http"
	"encoding/json"

	"github.com/anirudh-devanand/Password-Manager/src/initialize"

	"go.mongodb.org/mongo-driver/bson"
	"github.com/gorilla/mux"
)

//Entry is the object that the password manager deals with it has a password field and a website field
type Entry struct {
	Website  string `json:"Website"`
	Password string `json:"Password"`
}


// Create adds an entry to the map with the given website and password
func Create(w http.ResponseWriter, r *http.Request) {

	var entry Entry

	//Create Decoder object from HTTP request body and then parse and decode object into entry as a String
	err := json.NewDecoder(r.Body).Decode(&entry)
	if err != nil {
        http.Error(w, "Failed to parse request body", http.StatusBadRequest)
		log.Println("Failed to parse POST request body")
        return
    }

	//Initialize EntryDB instance
	entryDB := initialize.GetDB()

	log.Println("Adding entry...")

	//Ensure validity of password and website
	if entry.Website == "" || entry.Password == "" {
		http.Error(w,"Invalid credentials entered", http.StatusBadRequest)
		log.Println("Invalid credentials entered")
		return
	} else {
		//Perform CREATE operation in EntryDB
		_, err := entryDB.InsertOne(context.TODO(), bson.M{"Website":entry.Website, "Password": entry.Password})
		if err!=nil{
			http.Error(w,"Error adding entry to database", http.StatusInternalServerError)
			log.Println("Entry not added: ",err)
			return
		}

		log.Println("Entry added")
	}
}


// Read retrieves and prints either the password for a given website entry
func Read(w http.ResponseWriter, r *http.Request) {

	//Extract route variables into a String[] array
	HTTPparameters := mux.Vars(r)

	var entry Entry
	entry.Website = HTTPparameters["website"]

	//Initialize EntryDB instance
	entryDB := initialize.GetDB()

	log.Println("Lookup requested...")

	var filter bson.M
	var result Entry

	//Ensure validity of website
	if entry.Website == "" { 
			http.Error(w,"Invalid credentials entered", http.StatusBadRequest)
			log.Println("Invalid credentials entered")
			return
		}

	//Set filter based on website
	filter = bson.M{"Website" : entry.Website}

	//Perform READ operation from EntryDB and decode BSON into String result
	err := entryDB.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		http.Error(w,"Error searching for entry in database", http.StatusInternalServerError)
		log.Println("Lookup failed: ",err)
		return
	}

	log.Println("Lookup successful")

	//Recode String into a JSON response
	jsonResponse, err := json.Marshal(result)
	if err != nil {
		http.Error(w, "Error creating JSON response", http.StatusInternalServerError)
		log.Println("Error creating JSON response: ", err)
		return
	}

	//Set appropriate HTTP response header
	w.Header().Set("Content-Type", "application/json")

	//Write HTTP response body
	_, err = w.Write(jsonResponse)
	if err != nil {
		http.Error(w, "Error writing JSON response", http.StatusInternalServerError)
		log.Println("Error creating JSON response: ",err)
		return
	}

}


// Update updates the password for a given website, and if the updated password is empty, calls del()
func Update(w http.ResponseWriter, r *http.Request) {

	var entry Entry

	err := json.NewDecoder(r.Body).Decode(&entry)
	if err != nil {
        http.Error(w, "Failed to parse request body", http.StatusBadRequest)
		log.Println("Failed to parse POST request body")
        return
    }

	//Initialize EntryDB instance
	entryDB := initialize.GetDB()
	
	log.Println("Updating entry...")

	//Ensure validity website
	if entry.Website == ""{
		http.Error(w,"Invalid credentials entered", http.StatusBadRequest)
		log.Println("Invalid credentials entered")
		return
	}

	//Route empty password update calls to DEL()
	if entry.Password == "" {
		Delete(w, r)
	}

	//Set filter based on website
	filter := bson.M{"Website" : entry.Website}
	//Set operator to update entry
	update := bson.M{"$set": bson.M{"Password" : entry.Password}}

	//Perform UPDATE operation in EntryDB
	_, err = entryDB.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		http.Error(w,"Error updating entry in database", http.StatusInternalServerError)
		log.Println("Update failed: ",err)
		return
	}

	log.Println("Update successful")
}


// Delete deletes an entry from the map for the given website
func Delete(w http.ResponseWriter, r *http.Request) {

	//Extract route variables
	HTTPparameters := mux.Vars(r)
	website := HTTPparameters["website"]

	//Initialize EntryDB instance
	entryDB := initialize.GetDB()

	log.Println("Deleting entry...")

	//Set filter based on website
	filter := bson.M{"Website": website}

	//Perform DELETE operation in EntryDB
	_, err := entryDB.DeleteOne(context.TODO(), filter)
	if err != nil {
		http.Error(w,"Error deleting entry in database", http.StatusInternalServerError)
		log.Println("Deletion unsuccessful: ",err)
		return
	}

	log.Println("Deletion successful")
}
