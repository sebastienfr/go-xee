package xee

import (
	"fmt"
)

const (
	userURL = "/v3/users/%v"
)

// GetMe return a user from an access token
func (s *SDK) GetMe(token string) (User, error) {
	uri := fmt.Sprintf(userURL, "me")
	user := User{}

	err := s.decodeGetRequest(uri, token, nil, &user)

	return user, err
}
