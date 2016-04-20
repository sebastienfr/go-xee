package xee

import (
	"time"
)

// Location struct
type Location struct {
	Latitude   float32   `json:"latitude"`
	Longitude  float32   `json:"longitude"`
	Altitude   float32   `json:"altitude"`
	Satellites int       `json:"satellites"`
	Date       time.Time `json:"date"`
}
