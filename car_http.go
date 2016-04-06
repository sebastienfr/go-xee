package xee

import (
    "fmt"
)

const (
    carsURL = "/v3/users/%d/cars"
    carURL  = "/v3/users/%d/cars/%d"
)

// FindCars for a given user
func (s *SDK) FindCars(userID int64, token string) ([]Car, error) {
    var cars = make([]Car, 0)
    var params = make(map[string]string)

    uri := fmt.Sprintf(carsURL, userID)
    err := s.decodeGetRequest(uri, token, params, &cars)

    return cars, err
}

// FindCarByID find with ID for a given user
func (s *SDK) FindCarByID(userID int64, carID int64, token string) (Car, error) {
    var car = Car{}
    var params = make(map[string]string)

    uri := fmt.Sprintf(carURL, userID, carID)
    err := s.decodeGetRequest(uri, token, params, &car)

    return car, err
}
