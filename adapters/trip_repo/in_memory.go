package trip_repo

import (
	"github.com/devlucky/maporable-api/models"
	"errors"
)

type TripRepo interface {
	List() ([]*models.Trip)
	Get(id string) (*models.Trip)
	Create(trip *models.Trip) (error)
}

func Test() (TripRepo) {
	return NewInMemory()
}


type InMemory struct {
	trips map[string]*models.Trip
}

func NewInMemory() (*InMemory) {
	return &InMemory{
		trips: make(map[string]*models.Trip),
	}
}

func (repo *InMemory) List() ([]*models.Trip) {
	list := make([]*models.Trip, 0, len(repo.trips))

	for  _, trip := range repo.trips {
		list = append(list, trip)
	}

	return list
}

func (repo *InMemory) Get(id string) (*models.Trip) {
	return repo.trips[id]
}

func (repo *InMemory) Create(trip *models.Trip) (error) {
	if _, ok := repo.trips[trip.Id]; ok {
		return errors.New("Duplicate ID for trip")
	}

	repo.trips[trip.Id] = trip
	return nil
}