package xee

import (
    "fmt"
)

const (
    userURL     = "/v1/user/%v.json"
)

// GetMe return a user from an access token
func (s *SDK) GetMe(token string) (User, error) {
	uri := fmt.Sprintf(s.host+userURL, "me")
	user := User{}

    err := s.decodeGetRequest(uri, token, nil, &user)

	return user, err
}
