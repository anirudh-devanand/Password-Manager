package main

import (
	"log"
	"net/http"

	"github.com/anirudh-devanand/PwdMngr-Go/src/initialize"
	"github.com/anirudh-devanand/PwdMngr-Go/src/models"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {

	//Initialze a router using mux
	router := mux.NewRouter()

	//Set REST API call paths and appropriate handlers
	router.HandleFunc("/entry", models.Create).Methods("POST")
	router.HandleFunc("/entry/{website}", models.Read).Methods("GET")
	router.HandleFunc("/entry/{website}", models.Update).Methods("PUT")
	router.HandleFunc("/entry/{website}", models.Delete).Methods("DELETE")

	//Initialize the first instance of EntryDB
	initialize.InitDB()

	//Implement middleware to allow CORS enables HTTP response with appropriate options
	c := cors.New(cors.Options{
        AllowedOrigins: []string{"*"},
        AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowedHeaders: []string{"*"},
        AllowCredentials: true,
    })

	// Use the CORS middleware with mux router
    handler := c.Handler(router)

	
	log.Println("Server is running on :8081")

	//Set router path root and instruct router to listen on localhost:8081
	http.Handle("/", router)
    http.ListenAndServe(":8081", handler)



}
