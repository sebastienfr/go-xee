package xee

import (
    "net/http"
    "net/url"
    "encoding/json"
    "fmt"
)

const (
    getMethod           = "GET"
    postMethod          = "POST"
    bearer              = "Bearer"
    authorizationHeader = "Authorization"
)

func (s *SDK) decodeGetRequest(u string, token string, params map[string]string, v interface{}) error {
    // Create URI
    uri, err := url.Parse(s.host)
	uri.Path += u

    // Add parameters
    parameters := url.Values{}
    for key, value := range params {
        parameters.Add(key, value)
    }

    uri.RawQuery = parameters.Encode()
	uriString := uri.String()

    // Prepare request
    req, err := http.NewRequest(getMethod, uriString, nil)

    if err != nil {
        return err
    }

    // Add token to headers
    req.Header.Add(authorizationHeader, fmt.Sprintf("Bearer %s", token))

    // Send request
    res, err := s.client.Do(req)
    defer res.Body.Close()

	if err != nil {
		return err
	} else if res.StatusCode == http.StatusForbidden {
		return ErrForbidden
    } else if res.StatusCode == http.StatusNotFound {
        return ErrEntityNotFound
	} else if res.StatusCode != http.StatusOK {
		return fmt.Errorf("GET %v on %v", res.StatusCode, uri)
	}

    err = json.NewDecoder(res.Body).Decode(v)

    return err
}
