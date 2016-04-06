package xee

import (
    "time"
)

// Signal struct
type Signal struct {
    Name   string    `json:"string"`
    Value  float32   `json:"string"`
    Source string    `json:"string"`
    Date   time.Time `json:"date"`
}
