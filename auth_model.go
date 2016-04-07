package xee

import (
	"time"
)

// TokenResponse struct
type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int    `json:"expires_in"`
	ReceivedDate time.Time
}
