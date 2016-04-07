package xee

import (
	"time"
)

// Car struct
type Car struct {
	ID             int64     `json:"id"`
	UUID           string    `json:"uuid"`
	Name           string    `json:"name"`
	Make           *string   `json:"make"`
	Year           int       `json:"year"`
	NumberPlate    *string   `json:"numberPlate"`
	DeviceID       *int64    `json:"deviceId"`
	CardbID        int64     `json:"cardbId"`
	CreationDate   time.Time `json:"creationDate"`
	LastUpdateDate time.Time `json:"lastUpdateDate"`
}
