package xee

import (
	"time"
)

// Trip struct
type Trip struct {
	ID            string    `json:"id"`
	BeginLocation Location  `json:"beginLocation"`
	EndLocation   Location  `json:"endLocation"`
	BeginDate     time.Time `json:"beginDate"`
	StopDate      time.Time `json:"stopDate"`
}
