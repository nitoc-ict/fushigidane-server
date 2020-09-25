package mapsapi

import (
	"os"

	"github.com/pkg/errors"
	"googlemaps.github.io/maps"
)

var (
	FushigidaneMaspApiClient *maps.Client
)

func InitMapsAPIClient() error {
	c, err := maps.NewClient(maps.WithAPIKey(os.Getenv("GMAPS_TOKEN")))
	if err != nil {
		return errors.Wrap(err, "failed make google maps api client")
	}

	FushigidaneMaspApiClient = c

	return nil
}
