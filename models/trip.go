package models

import (
	"github.com/satori/go.uuid"
)

// A trip represents a one-time visit to a particular country
type Trip struct {
	Id string `json:"id"`
	User string `json:"user"`
	Country string `json:"country"`
	Status string `json:"status"`
	StartDate string `json:"start_date"`
	EndDate string `json:"end_date"`
	Latitude float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

func NewTrip(user, country, status, startDate, endDate string) (*Trip, error) {
	err := validateDate(startDate)
	if err != nil {
		return nil, err
	}

	err = validateDate(endDate)
	if err != nil {
		return nil, err
	}

	trip := &Trip{
		Id: uuid.NewV4().String(),
		User: user,
		Country: country,
		Status: status,
		StartDate: startDate,
		EndDate: endDate,
		Latitude: 40.40,
		Longitude: 50.50,
	}

	return trip, nil
}

// TODO: Actual functionality
func getCoords(country string) (lat float64, long float64) {
	lat, long = 40.40, 50.50
	return
}
