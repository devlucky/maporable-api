package config

import (
	"github.com/devlucky/maporable-api/adapters/trip_repo"
	"golang.org/x/oauth2"
)

type Config struct {
	TripRepo           trip_repo.TripRepo

	// Facebook OAuth
	FacebookOAuth      *oauth2.Config
	FacebookOAuthState string
}
