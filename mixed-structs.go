package main

import (
	"os"
)

// START CLIENT OMIT
type APIClient struct {
	apiAddress string
}

func NewAPIClient(addr string) *APIClient {
	return &APIClient{
		apiAddress: addr,
	}
}

// END CLIENT OMIT

func bad() {

	// START SERVER OMIT
	type Application struct {
		APIAddress string
		Client     *APIClient `json:"-"`
	}

	app := &Application{
		APIAddress: os.Getenv("API_ADDRESS"),
	}

	client := NewAPIClient(app.APIAddress)
	app.Client = client

	// END SERVER OMIT
}

func good() {

	// START GOOD_SERVER OMIT
	type ApplicationConfig struct {
		APIAddress string
	}

	appCfg := &ApplicationConfig{
		APIAddress: os.Getenv("API_ADDRESS"),
	}

	type Application struct {
		Client *APIClient
	}

	app := &Application{
		Client: NewAPIClient(appCfg.APIAddress),
	}

	// END GOOD_SERVER OMIT
}
