package gettransitpoints

import (
	"math"

	"github.com/nitoc-ict/fushigidane-server/convertaddress"
	"github.com/nitoc-ict/fushigidane-server/rdbms"
	"github.com/pkg/errors"
)

func PullTransitPoints(label string) ([]TransitPoint, error) {
	var transitpoints []TransitPoint

	rows, err := rdbms.Fushigidane.Query(`SELECT * FROM transitpoints WHERE (label = ?);`, label)
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

func SearchCandidatePoint(origin, destination, label string) ([]TransitPoint, error) {
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

	return transitpoints, nil
}

func euclideanDistance(originLat, originLong, destinationLat, destinationLong float64) float64 {
	return math.Sqrt(((destinationLat - originLat) * (destinationLat - originLat)) +
		((destinationLong - destinationLong) * (destinationLong - destinationLong)))

}
