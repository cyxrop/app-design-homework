package room

import "errors"

var (
	ErrHotelIsNotFound    = errors.New("hotel is not found")
	ErrRoomIsNotFound     = errors.New("room is not found")
	ErrRoomIsNotAvailable = errors.New("room is not available")
)
