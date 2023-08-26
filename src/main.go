package main

import(
	"log"
	"net/http"

	"github.com/anirudh-devanand/PwdMngr-Go/src/models"
	"github.com/anirudh-devanand/PwdMngr-Go/src/initialize"

	"github.com/gorilla/mux"
)






func main(){ 

	
	router := mux.NewRouter()

    router.HandleFunc("/entry", models.Create).Methods("POST")
	router.HandleFunc("/entry/{website}", models.Read).Methods("GET")
	router.HandleFunc("/entry/{website}", models.Update).Methods("PUT")
	router.HandleFunc("/entry/{website}", models.Delete).Methods("DELETE")

	initialize.InitDB()

    http.Handle("/", router)
    log.Println("Server is running on :8080")
    http.ListenAndServe(":8080", nil)
}
