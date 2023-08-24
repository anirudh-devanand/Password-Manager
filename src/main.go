package main

import(
	"fmt"
	"net/http"

	"github.com/anirudh-devanand/PwdMngr-Go/src/models"

	"github.com/gorilla/mux"
)






func main(){ 

	
	router := mux.NewRouter()

    router.HandleFunc("/entry", models.Create).Methods("POST")
	router.HandleFunc("/entry/{website}", models.Read).Methods("GET")
	router.HandleFunc("/entry/{website}", models.Update).Methods("PUT")
	router.HandleFunc("/entry/{website}", models.Delete).Methods("DELETE")

    http.Handle("/", router)
    fmt.Println("Server is running on :8080")
    http.ListenAndServe(":8080", nil)
}
