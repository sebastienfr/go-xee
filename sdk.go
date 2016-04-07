package xee

import (
	"net/http"
	"time"
)

const (
	cloud   = "https://cloud.xee.com"
	sandbox = "https://staging.xee.com"
	timeout = 10
)

// SDK struct is used to handle API calls
type SDK struct {
	host         string
	redirect     string
	clientID     string
	clientSecret string
	client       *http.Client
}

// NewSDK create an new instance of SDK
func NewSDK(identifier string, secret string, redirect string) *SDK {
	return &SDK{
		host:         cloud,
		redirect:     redirect,
		clientID:     identifier,
		clientSecret: secret,
		client:       &http.Client{Timeout: timeout * time.Second},
	}
}

// SetSandbox to use pre-production environnment
func (sdk *SDK) SetSandbox(isSandbox bool) {
	if isSandbox {
		sdk.host = sandbox
	} else {
		sdk.host = cloud
	}
}
