package order

import (
	"applicationDesignTest/internal/entity/date"
	"applicationDesignTest/internal/entity/room"
	"errors"
	"time"
)

var (
	ErrInvalidOrder = errors.New("invalid order")
)

type Order struct {
	HotelID   string
	RoomID    string
	UserEmail string
	Interval  date.Interval
}

func (o Order) Validate() error {
	if o.HotelID == "" {
		return errors.New("empty hotel id")
	}
	if o.RoomID == "" {
		return errors.New("empty room id")
	}
	if o.UserEmail == "" {
		return errors.New("empty user email")
	}
	if o.Interval.From.IsZero() {
		return errors.New("from is empty")
	}
	if o.Interval.To.IsZero() {
		return errors.New("to is empty")
	}
	if o.Interval.From.After(o.Interval.To) {
		return errors.New("from is after to")
	}
	return nil
}

func (o Order) TryReserve(availabilities room.Availabilities) []time.Time {
	daysToBook := o.Interval.Days()

	unavailableDaysMap := make(map[time.Time]struct{})
	for _, day := range daysToBook {
		unavailableDaysMap[day] = struct{}{}
	}

	for _, dayToBook := range daysToBook {
		for i, availability := range availabilities {
			var (
				shouldReserve = o.HotelID == availability.HotelID &&
					o.RoomID == availability.RoomID &&
					availability.Date.Equal(dayToBook)

				hasQuota = availability.Quota > 0
			)

			if !shouldReserve || !hasQuota {
				continue
			}

			availability.Quota -= 1
			availabilities[i] = availability
			delete(unavailableDaysMap, dayToBook)
		}
	}

	unavailableDays := make([]time.Time, 0, len(unavailableDaysMap))
	for day := range unavailableDaysMap {
		unavailableDays = append(unavailableDays, day)
	}
	return unavailableDays
}
