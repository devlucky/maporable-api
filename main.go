package main  // import "github.com/devlucky/maporable-api"

import (
	"fmt"
	"net/http"
	"log"
	"github.com/julienschmidt/httprouter"
	"github.com/devlucky/maporable-api/api"
	"github.com/devlucky/maporable-api/adapters/trip_repo"
	"github.com/devlucky/maporable-api/config"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/facebook"
	"encoding/json"
	"io/ioutil"
)

const userConfFilename string = "conf.json"

// TODO: Move this to API
func Ping(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprint(w, "pong")
}

func InjectConfig(a *config.Config, f func (http.ResponseWriter, *http.Request, httprouter.Params, *config.Config)) (func (http.ResponseWriter, *http.Request, httprouter.Params)) {
	return func (w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		f(w, r, ps, a)
	}
}

type UserConfig struct {
	FbClientID string `json:"_fb_client_id"`
	FbClientSecret string `json:"fb_client_secret"`
	FbRedirectURL string `json:"_fb_redirect_url"`
	FbState string `json:"_fb_state "`
}

func GetConfigVars() (*UserConfig) {
	var uConf UserConfig
	conf, err := ioutil.ReadFile(userConfFilename)
	if err != nil {
		log.Fatalf("Error reading from file %s", userConfFilename)
	}

	err = json.Unmarshal(conf, uConf)
	if err != nil {
		log.Fatalf("Error unmarshaling conf in %s", userConfFilename)
	}

	return uConf
}

func CurrentConfig(uConf *UserConfig) (*config.Config) {
	return &config.Config{
		TripRepo: trip_repo.NewInMemory(),
		FacebookOAuth: &oauth2.Config{
			ClientID:     uConf.FbClientID,
			ClientSecret: uConf.FbClientSecret,
			RedirectURL:  uConf.FbRedirectURL,
			Scopes:       []string{"public_profile"},
			Endpoint:     facebook.Endpoint,
		},
		FacebookOAuthState: uCong.FbState,
	}
}

func main() {
	conf := CurrentConfig(GetConfigVars())
	router := httprouter.New()

	// Ping-pong
	router.GET("/", Ping)

	// Authentication endpoints
	router.POST("/login", InjectConfig(conf, api.Login))
	router.POST("/login/facebook", InjectConfig(conf, api.LoginWithFacebook))

	// Trips endpoints
	router.GET("/trips", InjectConfig(conf, api.GetTripsList))
	router.POST("/trips", InjectConfig(conf, api.CreateTrip))
	router.GET("/trips/:id", InjectConfig(conf, api.GetTrip))

	log.Println("Listening on 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

