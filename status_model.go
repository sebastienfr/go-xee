package xee

import "time"

// Status struct
type Status struct {
	Signals        []Signal      `json:"signals"`
	CreationDate   time.Time     `json:"creationDate"`
	LastUpdateDate time.Time     `json:"lastUpdateDate"`
	Location       Location      `json:"location"`
	Accelerometer  Accelerometer `json:"accelerometer"`
}

// Accelerometer struct
type Accelerometer struct {
	X    float64   `json:"x"`
	Y    float64   `json:"y"`
	Z    float64   `json:"z"`
	Date time.Time `json:"date"`
}
