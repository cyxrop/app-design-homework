package date

import (
	"time"
)

type Date struct {
	Year  int
	Month int
	Day   int
}

type Interval struct {
	From time.Time
	To   time.Time
}

func (i Interval) Days() []time.Time {
	days := make([]time.Time, 0)
	for day := i.From; !day.After(i.To); day = day.AddDate(0, 0, 1) {
		days = append(days, day)
	}
	return days
}

func (i Interval) Has(date time.Time) bool {
	return (i.From.Before(date) || i.From.Equal(date)) && (i.To.Equal(i.To) || i.To.After(date))
}
