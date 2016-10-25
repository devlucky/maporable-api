package api

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"github.com/devlucky/maporable-api/models"
	"encoding/json"
	"log"
	"fmt"
	"github.com/devlucky/maporable-api/config"
	"net/url"
	"strings"
	"golang.org/x/oauth2"
	"io/ioutil"
)

func Login(w http.ResponseWriter, r *http.Request, ps httprouter.Params, a *config.Config) {
	Url, err := url.Parse(a.FacebookOAuth.Endpoint.AuthURL)
	if err != nil {
		log.Fatal("Parse: ", err)
	}
	parameters := url.Values{}
	parameters.Add("client_id", a.FacebookOAuth.ClientID)
	parameters.Add("scope", strings.Join(a.FacebookOAuth.Scopes, " "))
	parameters.Add("redirect_uri", a.FacebookOAuth.RedirectURL)
	parameters.Add("response_type", "code")
	parameters.Add("state", a.FacebookOAuthState)

	Url.RawQuery = parameters.Encode()
	http.Redirect(w, r, Url.String(), http.StatusTemporaryRedirect)
}

func LoginWithFacebook(w http.ResponseWriter, r *http.Request, ps httprouter.Params, a *config.Config) {
	state := r.FormValue("state")
	if state != a.FacebookOAuthState {
		log.Printf("invalid oauth state, expected '%s', got '%s'\n", a.FacebookOAuthState, state)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	code := r.FormValue("code")
	token, err := a.FacebookOAuth.Exchange(oauth2.NoContext, code)
	if err != nil {
		fmt.Printf("oauthConf.Exchange() failed with '%s'\n", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	resp, err := http.Get("https://graph.facebook.com/me?access_token=" +
		url.QueryEscape(token.AccessToken))
	if err != nil {
		fmt.Printf("Get: %s\n", err)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}
	defer resp.Body.Close()

	response, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("ReadAll: %s\n", err)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	log.Printf("parseResponseBody: %s\n", string(response))

	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
}
