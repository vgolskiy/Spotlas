package main

import (
	"endpoint/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/spots/id/{id}", handlers.GetSpotByID)
	router.HandleFunc("/spots", handlers.GetSpots)
	router.HandleFunc("/spots/parameters", handlers.GetSpotByParameters)
	log.Fatal(http.ListenAndServe(":8000", router))
}
