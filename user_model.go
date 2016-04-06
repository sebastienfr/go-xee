package xee

import(
    "time"
)

// User struct
type User struct {
    ID                int64     `json:"id"`
    UUID              string    `json:"uuid"`
    LastName          string    `json:"lastName"`
    FirstName         string    `json:"firstName"`
    Nickname          *string   `json:"nickname"`
    Gender            string    `json:"gender"`
    Role              string    `json:"role"`
    IsLocationEnabled bool      `json:"isLocationEnabled"`
    CreationDate      time.Time `json:"creationDate"`
    LastUpdateDate    time.Time `json:"lastUpdateDate"`
}
