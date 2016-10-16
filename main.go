package main  // import "github.com/devlucky/maporable-api"

import (
	"fmt"
	"net/http"
	"log"
	"github.com/julienschmidt/httprouter"
	"github.com/devlucky/maporable-api/api"
)

func Ping(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprint(w, "pong")
}

func main() {
	router := httprouter.New()
	router.GET("/", Ping)

	router.POST("/trips", api.CreateTrip)

	log.Println("Listening on 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

