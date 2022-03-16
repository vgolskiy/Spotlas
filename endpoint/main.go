package main

import (
	"endpoint/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/spots/{id}", handlers.GetSpotByID)
	router.HandleFunc("/spots", handlers.GetSpots)
	log.Fatal(http.ListenAndServe(":8000", router))
}
