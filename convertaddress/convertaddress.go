package convertaddress

import (
	"context"
	"os"

	"github.com/pkg/errors"
	"googlemaps.github.io/maps"
)

func GetLonLat(address string) (Coordinate, error) {
	var coordinate Coordinate
	c, err := maps.NewClient(maps.WithAPIKey(os.Getenv("MAPS_APIKEY")))
	if err != nil {
		return coordinate, errors.Wrap(err, "Failed make New client")
	}

	r := &maps.GeocodingRequest{
		Address: address,
	}

	route, err := c.Geocode(context.Background(), r)
	if err != nil {
		return coordinate, errors.Wrap(err, "Failed geocode")
	}

	coordinate = Coordinate{
		Address:   address,
		Latitude:  route[0].Geometry.Location.Lat,
		Longitude: route[0].Geometry.Location.Lng,
	}

	return coordinate, nil
}
