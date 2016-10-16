package models

import "errors"

// A trip represents a one-time visit to a particular country
type Trip struct {
	Id string `json:"id"`
	User string `json:"user"`
	Place string `json:"place"`
	Latitude float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Description string `json:"description"`
}

func NewTrip(place string) (*Trip, error) {
	if place == "" {
		return nil, errors.New("The place cannot be empty")
	}

	trip := &Trip{
		User: "hector",
		Place: place,
		Latitude: 40.40,
		Longitude: 50.50,
		Description: "lots of sharks and groundhogs",
	}
	return trip, nil
}
