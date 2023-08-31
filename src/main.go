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

	router := mux.NewRouter()

	router.HandleFunc("/entry", models.Create).Methods("POST")
	router.HandleFunc("/entry/{website}", models.Read).Methods("GET")
	router.HandleFunc("/entry/{website}", models.Update).Methods("PUT")
	router.HandleFunc("/entry/{website}", models.Delete).Methods("DELETE")

	initialize.InitDB()

	c := cors.New(cors.Options{
        AllowedOrigins: []string{"*"},
        AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowedHeaders: []string{"*"},
        AllowCredentials: true,
    })


    handler := c.Handler(router)

	
	log.Println("Server is running on :8081")

	http.Handle("/", router)
    http.ListenAndServe(":8081", handler)



}
