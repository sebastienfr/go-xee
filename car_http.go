package xee

import (
    "fmt"
)

const (
    carsURL = "/v3/users/%d/cars"
)

// FindCars for a given user
func (s *SDK) FindCars(userID int64, token string) ([]Car, error) {
    var cars = make([]Car, 0)
    var params = make(map[string]string)

    uri := fmt.Sprintf(carsURL, userID)
    err := s.decodeGetRequest(uri, token, params, &cars)

    return cars, err
}
