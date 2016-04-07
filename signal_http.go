package xee

import (
	"fmt"
	"time"
)

const (
	signalsURL       = "/v3/cars/%d/signals"
	signalsByTripURL = "/v3/cars/%d/trips/%s/signals"
)

// FindSignals for a given car
func (s *SDK) FindSignals(carID int64, token string, names *[]string, limit *int, begin *time.Time, stop *time.Time) ([]Location, error) {
	var locations = make([]Location, 0)
	var params = make(map[string]string)

	if limit != nil {
		params["limit"] = fmt.Sprintf("%v", *limit)
	}

	if begin != nil {
		params["begin"] = (*begin).String()
	}

	if stop != nil {
		params["stop"] = (*stop).String()
	}

	if names != nil {
		/*for _, name := range *names {
		  }*/
	}

	uri := fmt.Sprintf(signalsURL, carID)
	err := s.decodeGetRequest(uri, token, params, &locations)

	return locations, err
}

// FindSignalsByTrip for a given car
func (s *SDK) FindSignalsByTrip(carID int64, tripID string, token string, names *[]string) ([]Location, error) {
	var locations = make([]Location, 0)
	var params = make(map[string]string)

	uri := fmt.Sprintf(signalsByTripURL, carID, tripID)
	err := s.decodeGetRequest(uri, token, params, &locations)

	return locations, err
}
