package room

import (
	"time"
)

type Availability struct {
	HotelID string
	RoomID  string
	Date    time.Time
	Quota   int
}

type Availabilities []Availability
