package gettransitpoints

import (
	"fmt"
	"math"
	"net/http"
	"strings"

	"github.com/kr/pretty"
	"github.com/labstack/echo"
	"github.com/nitoc-ict/fushigidane-server/convertaddress"
	"github.com/nitoc-ict/fushigidane-server/rdbms"
	"github.com/pkg/errors"
)

func GetTransitPoints(c echo.Context) error {
	var routeData RouteRequest

	err := c.Bind(&routeData)

	if err != nil {
		c.JSON(http.StatusBadRequest, `{"status": "error bind json"}`)

		return nil
	}

	if !strings.Contains(routeData.Origin, "沖縄") || !strings.Contains(routeData.Origin, "okinawa") {
		c.JSON(http.StatusBadRequest, `{"status": "error please address for okinawa"}`)

		return nil
	}

	if !strings.Contains(routeData.Destination, "沖縄") || !strings.Contains(routeData.Destination, "okinawa") {
		c.JSON(http.StatusBadRequest, `{"status": "error please address for okinawa"}`)

		return nil
	}

	transitpoints, err := SearchCandidatePoint(routeData.Origin, routeData.Destination, routeData.Scenes)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, `{"status": "error server"}`)
		return nil
	}

	c.JSON(http.StatusOK, transitpoints)

	return nil
}

func PullTransitPoints(label []string) ([]TransitPoint, error) {
	var transitpoints []TransitPoint

	if len(label) == 1 {
		rows, err := rdbms.Fushigidane.Query(`SELECT * FROM transitpoints WHERE (label = ?);`, label[0])
		if err != nil {
			return nil, errors.Wrap(err, "Failed Query `SELECT * FROM transitpoints WHERE label`")
		}

		for rows.Next() {
			transitpoint := TransitPoint{}
			err = rows.Scan(&transitpoint.Id, &transitpoint.Address, &transitpoint.Label, &transitpoint.Latitude, &transitpoint.Longitude)

			if err != nil {
				return nil, errors.Wrap(err, "Failed Scan transitpoint")
			}

			transitpoints = append(transitpoints, transitpoint)
		}
	} else {
		for _, e := range label {
			rows, err := rdbms.Fushigidane.Query(`SELECT * FROM transitpoints WHERE (label = ?);`, e)
			if err != nil {
				return nil, errors.Wrap(err, "Failed Query `SELECT * FROM transitpoints WHERE label`")
			}

			for rows.Next() {
				transitpoint := TransitPoint{}
				err = rows.Scan(&transitpoint.Id, &transitpoint.Address, &transitpoint.Label, &transitpoint.Latitude, &transitpoint.Longitude)

				if err != nil {
					return nil, errors.Wrap(err, "Failed Scan transitpoint")
				}

				transitpoints = append(transitpoints, transitpoint)
			}
		}
	}

	return transitpoints, nil
}

func SearchCandidatePoint(origin, destination string, label []string) ([]TransitPoint, error) {
	var transitpoints []TransitPoint

	candidatePoint, err := PullTransitPoints(label)
	if err != nil {
		return nil, errors.Wrap(err, "failed pull transit points")
	}

	originPoint, err := convertaddress.GetLonLat(origin)
	if err != nil {
		return nil, errors.Wrap(err, "failed get origin point lon lat")
	}
	destinationPoint, err := convertaddress.GetLonLat(destination)
	if err != nil {
		return nil, errors.Wrap(err, "failed get destination point lon lat")
	}

	transitpoints = append(transitpoints, TransitPoint{Id: 0, Address: origin, Latitude: originPoint.Latitude, Longitude: originPoint.Longitude})

	distanceToDestination := euclideanDistance(originPoint.Latitude, originPoint.Longitude, destinationPoint.Latitude, destinationPoint.Longitude)

	next := candidatePoint[0]
	minDisToOrigin := euclideanDistance(candidatePoint[0].Latitude, candidatePoint[0].Longitude, originPoint.Latitude, originPoint.Longitude)

	for _, e := range candidatePoint {
		tmpDisToDestination := euclideanDistance(e.Latitude, e.Longitude, destinationPoint.Latitude, destinationPoint.Longitude)
		tmpDisToOrigin := euclideanDistance(e.Latitude, e.Longitude, originPoint.Latitude, originPoint.Longitude)

		if tmpDisToOrigin < minDisToOrigin && tmpDisToDestination < distanceToDestination {
			next = e
			distanceToDestination = tmpDisToDestination
			minDisToOrigin = tmpDisToOrigin
		}
	}
	transitpoints = append(transitpoints, next)

	for {
		distanceToDestination = euclideanDistance(transitpoints[len(transitpoints)-1].Latitude, transitpoints[len(transitpoints)-1].Longitude, destinationPoint.Latitude, destinationPoint.Longitude)
		minDisToOrigin = euclideanDistance(candidatePoint[0].Latitude, candidatePoint[0].Longitude, transitpoints[len(transitpoints)-1].Latitude, transitpoints[len(transitpoints)-1].Longitude)
		tmp := next
		for _, e := range candidatePoint {
			tmpDisToDestination := euclideanDistance(e.Latitude, e.Longitude, destinationPoint.Latitude, destinationPoint.Longitude)
			tmpDisToOrigin := euclideanDistance(e.Latitude, e.Longitude, transitpoints[len(transitpoints)-1].Latitude, transitpoints[len(transitpoints)-1].Longitude)

			if tmpDisToOrigin < minDisToOrigin && tmpDisToDestination < distanceToDestination {
				next = e
				distanceToDestination = tmpDisToDestination
				minDisToOrigin = tmpDisToOrigin
			}
		}

		if tmp == next {
			break
		}

		transitpoints = append(transitpoints, next)
	}

	var waypoints []string

	for _, e := range transitpoints {
		waypoints = append(waypoints, e.Address)
	}

//	r := &maps.DirectionsRequest{
//		Origin:      origin,
//		Destination: destination,
//		Waypoints:   waypoints,
//	}
//
//	route, _, err := mapsapi.FushigidaneMaspApiClient.Directions(context.Background(), r)

//	pretty.Println(route)

	pretty.Println(transitpoints)


	transitpoints = append(transitpoints, TransitPoint{Id: 0, Address: origin, Latitude: destinationPoint.Latitude, Longitude: destinationPoint.Longitude})

	return transitpoints, nil
}

func euclideanDistance(originLat, originLong, destinationLat, destinationLong float64) float64 {
	return math.Sqrt(((destinationLat - originLat) * (destinationLat - originLat)) +
		((destinationLong - destinationLong) * (destinationLong - destinationLong)))

}
