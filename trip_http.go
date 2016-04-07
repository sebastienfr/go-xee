package xee

import (
	"fmt"
	"time"
)

const (
	tripURL = "/v3/cars/%d/trips"
)

// FindTrips for a given car
func (s *SDK) FindTrips(carID int64, token string, begin *time.Time, stop *time.Time) ([]Trip, error) {
	var trips = make([]Trip, 0)
	var params = make(map[string]string)

	if begin != nil {
		params["begin"] = (*begin).String()
	}

	if stop != nil {
		params["stop"] = (*stop).String()
	}

	uri := fmt.Sprintf(tripURL, carID)
	err := s.decodeGetRequest(uri, token, params, &trips)

	return trips, err
}
