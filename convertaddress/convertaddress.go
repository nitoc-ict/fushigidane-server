package convertaddress

import (
	"context"
	"net/http"
	"os"

	"github.com/labstack/echo"
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
