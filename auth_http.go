package xee

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"
)

const (
	authURL  = "/v1/auth/auth"
	tokenURL = "/v1/auth/access_token.json"
)

// GetAuthURI to allow user to connect
func (s *SDK) GetAuthURI(state string) string {
	u, _ := url.Parse(s.host + authURL)

	parameters := url.Values{}
	parameters.Add("client_id", s.clientID)
	parameters.Add("redirect_uri", s.redirect)
	parameters.Add("state", state)
	u.RawQuery = parameters.Encode()

	return u.String()
}

// GetTokenFromCode is used to get a Token from a code
func (s *SDK) GetTokenFromCode(code string) (TokenResponse, error) {
	return s.getToken(code, false)
}

// GetTokenFromRefreshToken is used to get a Token from a code
func (s *SDK) GetTokenFromRefreshToken(refreshToken string) (TokenResponse, error) {
	return s.getToken(refreshToken, true)
}

func (s *SDK) getToken(code string, refresh bool) (TokenResponse, error) {
	t := TokenResponse{}

	data := url.Values{}
	data.Add("redirect_uri", s.redirect)

	if refresh {
		data.Add("grant_type", "refresh_token")
		data.Add("refresh_token", code)
	} else {
		data.Add("grant_type", "authorization_code")
		data.Add("code", code)
	}

	req, _ := http.NewRequest(postMethod, s.host+tokenURL, bytes.NewBufferString(data.Encode()))
	req.SetBasicAuth(s.clientID, s.clientSecret)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

	res, err := s.client.Do(req)

	if err != nil {
		return t, err
	}

	defer res.Body.Close()
	json.NewDecoder(res.Body).Decode(&t)

	return t, err
}
