package api

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"github.com/devlucky/maporable-api/models"
	"encoding/json"
	"log"
	"fmt"
)

type CreateTripInput struct {
	Place string `json:"place"`
}

func CreateTrip(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var input CreateTripInput

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&input)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid request body. Could not be parsed into JSON"))
		return
	}
	defer r.Body.Close()

	trip, err := models.NewTrip(input.Place)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		msg := fmt.Sprintf("Invalid trip parameter: %s", err.Error())
		w.Write([]byte(msg))
		return
	}


	// TODO: Save it


	jsonTrip, err := json.Marshal(trip)
	if err != nil {
		log.Printf("Unexpected error %s when marshaling the trip into JSON", err.Error())
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(jsonTrip))
}