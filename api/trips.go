package api

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"github.com/devlucky/maporable-api/models"
	"encoding/json"
	"log"
	"fmt"
	"github.com/devlucky/maporable-api/config"
)

func CreateTrip(w http.ResponseWriter, r *http.Request, ps httprouter.Params, a *config.Config) {
	var input models.Trip

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&input)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid request body. Could not be parsed into JSON"))
		return
	}
	defer r.Body.Close()

	trip, err := models.NewTrip(input.User, input.Country, input.Status, input.StartDate, input.EndDate)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		msg := fmt.Sprintf("Invalid trip parameter: %s", err.Error())
		w.Write([]byte(msg))
		return
	}

	err = a.TripRepo.Create(trip)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	jsonTrip, err := json.Marshal(trip)
	if err != nil {
		log.Printf("Unexpected error %s when marshaling the trip into JSON", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(jsonTrip))
}

func GetTripsList(w http.ResponseWriter, r *http.Request, ps httprouter.Params, a *config.Config) {
	trips := a.TripRepo.List()
	jsonTrips, err := json.Marshal(trips)
	if err != nil {
		log.Printf("Unexpected error %s when marshaling the trip into JSON", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(jsonTrips))
}


func GetTrip(w http.ResponseWriter, r *http.Request, ps httprouter.Params, a *config.Config) {
	id := ps.ByName("id")

	trip := a.TripRepo.Get(id)
	if trip == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	jsonTrip, err := json.Marshal(trip)
	if err != nil {
		log.Printf("Unexpected error %s when marshaling the trip into JSON", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(jsonTrip))
}