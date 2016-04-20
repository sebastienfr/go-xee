package xee

import (
	"time"
)

// Signal struct
type Signal struct {
	Date  time.Time `json:"date"`
	Value float64   `json:"value"`
	Name  string    `json:"name"`
}
