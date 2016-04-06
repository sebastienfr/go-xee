package xee

import (
    "fmt"
)

const (
    carStatusURL = "/v3/cars/%d/status"
)

// FindCarStatus for a given car
func (s *SDK) FindCarStatus(carID int64, token string) (Status, error) {
    var status = Status{}
    var params = make(map[string]string)

    uri := fmt.Sprintf(carStatusURL, carID)
    err := s.decodeGetRequest(uri, token, params, &status)

    return status, err
}
