package handlers

import (
	"encoding/json"
	"endpoint/enteties"
	"endpoint/repositary"
	"endpoint/utils"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func GetSpots(w http.ResponseWriter, r *http.Request) {
	var parameters enteties.Parameters

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&parameters)
	if err != nil || !utils.IsValidType(parameters.Type) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	log.Printf("Getting spots list for %v\n", parameters)
	err, spots := repositary.GetSpotsByParameters(&parameters)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(spots)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func GetSpotByID(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	log.Printf("Getting spot with id %v\n", id)
	err, spot := repositary.GetSpotByID(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(spot)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}