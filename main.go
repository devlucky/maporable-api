package main  // import "github.com/devlucky/maporable-api"

import (
	"fmt"
	"net/http"
	"log"
	"github.com/julienschmidt/httprouter"
	"github.com/devlucky/maporable-api/api"
	"github.com/devlucky/maporable-api/adapters/trip_repo"
	"github.com/devlucky/maporable-api/config"
)

// TODO: Move this to API
func Ping(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprint(w, "pong")
}

func InjectConfig(a *config.Adapters, f func (http.ResponseWriter, *http.Request, httprouter.Params, *config.Adapters)) (func (http.ResponseWriter, *http.Request, httprouter.Params)) {
	return func (w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		f(w, r, ps, a)
	}
}

func Config() (*config.Adapters) {
	return &config.Adapters{
		TripRepo: trip_repo.NewInMemory(),
	}
}

func main() {
	config := Config()
	router := httprouter.New()
	router.GET("/", Ping)

	router.GET("/trips", InjectConfig(config, api.GetTripsList))
	router.POST("/trips", InjectConfig(config, api.CreateTrip))
	router.GET("/trips/:id", InjectConfig(config, api.GetTrip))

	log.Println("Listening on 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

