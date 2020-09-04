package gettransitpoints

import (
	"github.com/nitoc-ict/fushigidane-server/rdbms"
	"github.com/pkg/errors"
)

func PullTransitPoints(label string) ([]TransitPoint, error) {
	var transitpoints []TransitPoint

	rows, err := rdbms.Fushigidane.Query(`SELECT * FROM transitpoints WHERE ?`, label)
	if err != nil {
		return nil, errors.Wrap(err, "Failed Query `SELECT * FROM transitpoints WHERE label`")
	}

	for rows.Next() {
		transitpoint := TransitPoint{}
		err = rows.Scan(&transitpoint.Id, &transitpoint.Address, &transitpoint.Label, &transitpoint.Latitude, &transitpoint.Longitude)

		if err != nil {
			return nil, errors.Wrap(err, "Failed Scan transitpoint")
		}
	}

	return transitpoints, nil
}
