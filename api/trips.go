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
		w.WriteHeader(http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(jsonTrip))
}


func GetTripsList(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	/*
	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("The page must be an integer"))
		return
	}

	pageSize, err := strconv.Atoi(r.URL.Query().Get("page_size"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("The page_size must be an integer"))
		return
	}
	*/

	// TODO: Get all trips from the database, filtered and paginated
	trip1, _ := models.NewTrip("sydney")
	trip2, _ := models.NewTrip("ottawa")
	trips := [2]*models.Trip{trip1, trip2}

	jsonTrips, err := json.Marshal(trips)
	if err != nil {
		log.Printf("Unexpected error %s when marshaling the trip into JSON", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(jsonTrips))
}


func GetTrip(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	fmt.Printf("id is %s", id)

	// TODO: Get from database
	trip, _ := models.NewTrip("sydney")
	jsonTrip, err := json.Marshal(trip)
	if err != nil {
		log.Printf("Unexpected error %s when marshaling the trip into JSON", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(jsonTrip))
}