package xee

import (
    "time"
)

// Trip struct
type Trip struct {
    UUID           string    `json:"uuid"`
    BeginLocation  Location  `json:"beginLocation"`
    StopLocation   Location  `json:"stopLocation"`
    BeginDate      time.Time `json:"beginDate"`
    StopDate       time.Time `json:"stopDate"`
    CreationDate   time.Time `json:"creationDate"`
    LastUpdateDate time.Time `json:"lastUpdateDate"`
}
