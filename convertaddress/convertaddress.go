package convertaddress

import (
	"context"
	"net/http"

	"github.com/labstack/echo"
	"github.com/nitoc-ict/fushigidane-server/mapsapi"
	"github.com/pkg/errors"
	"googlemaps.github.io/maps"
)

func ConvertAddress(c echo.Context) error {
	address := c.QueryParam("address")

	coordinate, err := GetLonLat(address)
	if err != nil {
		if err.Error() == "Failed make New client" {
			c.JSON(http.StatusBadRequest, `{"status": "status bad request"}`)
			return nil
		}
	}

	c.JSON(http.StatusOK, coordinate)

	return nil
}

func GetLonLat(address string) (Coordinate, error) {
	var coordinate Coordinate
	r := &maps.GeocodingRequest{
		Address: address,
	}

	route, err := mapsapi.FushigidaneMaspApiClient.Geocode(context.Background(), r)
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
