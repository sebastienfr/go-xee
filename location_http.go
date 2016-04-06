package xee

import (
    "fmt"
    "time"
)

const (
    locationsURL       = "/cars/%d/locations"
    locationsByTripURL = "/cars/{carId}/trips/%s/locations"
)

// FindLocations for a given car
func (s *SDK) FindLocations(carID int64, token string, limit *int, begin *time.Time, stop *time.Time) ([]Location, error) {
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

    uri := fmt.Sprintf(locationsURL, carID)
    err := s.decodeGetRequest(uri, token, params, &locations)

    return locations, err
}


// FindLocationsByTrip for a given car
func (s *SDK) FindLocationsByTrip(tripID string, token string) ([]Location, error) {
    var locations = make([]Location, 0)
    var params = make(map[string]string)

    uri := fmt.Sprintf(locationsByTripURL, tripID)
    err := s.decodeGetRequest(uri, token, params, &locations)

    return locations, err
}
