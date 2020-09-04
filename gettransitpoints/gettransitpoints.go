package gettransitpoints

import (
	"log"
	"math"

	"github.com/kr/pretty"
	"github.com/nitoc-ict/fushigidane-server/convertaddress"
	"github.com/nitoc-ict/fushigidane-server/rdbms"
	"github.com/pkg/errors"
)

func PullTransitPoints(label string) ([]TransitPoint, error) {
	var transitpoints []TransitPoint

	dbprepare, err := rdbms.Fushigidane.Prepare(`SELECT * FROM transitpoints WHERE label = 'sea';`)
	if err != nil {
		return nil, errors.Wrap(err, "Failed make prepare")
	}
	defer dbprepare.Close()

	rows, err := dbprepare.Query()
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

	return transitpoints, nil
}

func SearchCandidatePoint(origin, destination, label string) []TransitPoint {
	var transitpoints []TransitPoint

	candidate, err := PullTransitPoints(label)
	if err != nil {
		log.Fatal(err)
	}

	originCoordinate, err := convertaddress.GetLonLat(origin)
	if err != nil {
		log.Println(err)
	}
	transitpoints = append(transitpoints, TransitPoint{
		Id:        0,
		Address:   originCoordinate.Address,
		Label:     label,
		Latitude:  originCoordinate.Latitude,
		Longitude: originCoordinate.Longitude,
	})

	destinationCoordinate, err := convertaddress.GetLonLat(destination)
	if err != nil {
		log.Println(err)
	}

	nearDestination := euclideanDistance(candidate[0].Latitude, candidate[0].Longitude, destinationCoordinate.Latitude, destinationCoordinate.Longitude)
	originPoint := TransitPoint{
		Id:        0,
		Address:   originCoordinate.Address,
		Label:     label,
		Latitude:  originCoordinate.Latitude,
		Longitude: originCoordinate.Longitude,
	}

	for _, e := range candidate {
		if euclideanDistance(e.Latitude, e.Longitude, destinationCoordinate.Latitude, destinationCoordinate.Longitude) < nearDestination {
			nearDestination = euclideanDistance(e.Latitude, e.Longitude, destinationCoordinate.Latitude, destinationCoordinate.Longitude)
		}
	}

	shortestOriginDistance := euclideanDistance(candidate[0].Latitude, candidate[0].Longitude, originCoordinate.Latitude, originCoordinate.Longitude)
	shortestDestinationDistance := euclideanDistance(candidate[0].Latitude, candidate[0].Longitude, destinationCoordinate.Latitude, destinationCoordinate.Longitude)
	shortestPoint := candidate[0]

	for shortestDestinationDistance != nearDestination {
		pretty.Println(originPoint)
		for _, e := range candidate {
			originDistance := euclideanDistance(e.Latitude, e.Longitude, originPoint.Latitude, originPoint.Longitude)
			destinationDistance := euclideanDistance(e.Latitude, e.Longitude, destinationCoordinate.Latitude, destinationCoordinate.Longitude)

			if originDistance+destinationDistance < shortestDestinationDistance+shortestOriginDistance {
				shortestDestinationDistance = destinationDistance
				shortestOriginDistance = originDistance
				shortestPoint = TransitPoint{
					Id:        e.Id,
					Address:   e.Address,
					Label:     e.Label,
					Latitude:  e.Latitude,
					Longitude: e.Longitude,
				}
			}
		}

		originPoint = shortestPoint
		transitpoints = append(transitpoints, shortestPoint)
	}

	transitpoints = append(transitpoints, TransitPoint{
		Id:        0,
		Address:   destinationCoordinate.Address,
		Label:     label,
		Latitude:  destinationCoordinate.Latitude,
		Longitude: destinationCoordinate.Longitude,
	})

	return transitpoints
}

func euclideanDistance(originLat, originLong, destinationLat, destinationLong float64) float64 {
	return math.Sqrt(((destinationLat - originLat) * (destinationLat - originLat)) +
		((destinationLong - destinationLong) * (destinationLong - destinationLong)))

}
