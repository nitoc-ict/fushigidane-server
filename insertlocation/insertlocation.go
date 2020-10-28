package insertlocation

import (
	"context"
	"log"

	"github.com/labstack/echo"
	"github.com/nitoc-ict/fushigidane-server/mapsapi"
	"github.com/nitoc-ict/fushigidane-server/rdbms"
	"github.com/pkg/errors"
	"googlemaps.github.io/maps"
)

func RecvLocationData(c echo.Context) error {
	var ld RecvData
	err := c.Bind(&ld)

	log.Println(ld)
	if err != nil {
		log.Println(err)
		return errors.Wrap(err, "failed bind location data")
	}

	for _, e := range ld.Spot {
		err := InsertLocation(e)

		if err != nil {
			return errors.Wrap(err, "failed insert location")
		}
	}

	return nil
}

func InsertLocation(ld LocationData) error {
	prepare, err := rdbms.Fushigidane.Prepare(`INSERT INTO transitpoints(address, label, latitude, longitude) VALUES (?, ?, ?, ?)`)
	if err != nil {
		return errors.Wrap(err, "failed make prepare")
	}

	if ld.Address == "" {
		ld.Address, err = ConvertCoordinateToAddress(ld.Latitude, ld.Longitude)
		if err != nil {
			return err
		}
	}

	if ld.Scenes == "" {
		ld.Scenes = "sea"
	}

	_, err = prepare.Exec(ld.Address, ld.Scenes, ld.Latitude, ld.Longitude)
	if err != nil {
		return errors.Wrap(err, "failed insert data")
	}

	return nil
}

func ConvertCoordinateToAddress(latitude, longitude float64) (string, error) {
	r := &maps.GeocodingRequest{
		LatLng: &maps.LatLng{
			Lat: latitude,
			Lng: longitude,
		},
	}

	address, err := mapsapi.FushigidaneMaspApiClient.Geocode(context.Background(), r)
	if err != nil {
		return "", errors.Wrap(err, "failed get address")
	}

	return address[0].FormattedAddress, nil
}
