package date

import (
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestInterval_Days(t *testing.T) {
	t.Parallel()

	cases := []struct {
		interval Interval
		expected []time.Time
	}{
		{
			interval: Interval{
				From: time.Date(2024, 8, 1, 0, 0, 0, 0, time.UTC),
				To:   time.Date(2024, 8, 3, 0, 0, 0, 0, time.UTC),
			},
			expected: []time.Time{
				time.Date(2024, 8, 1, 0, 0, 0, 0, time.UTC),
				time.Date(2024, 8, 2, 0, 0, 0, 0, time.UTC),
				time.Date(2024, 8, 3, 0, 0, 0, 0, time.UTC),
			},
		},
		{
			interval: Interval{
				From: time.Date(2024, 8, 1, 0, 0, 0, 0, time.UTC),
				To:   time.Date(2024, 7, 30, 0, 0, 0, 0, time.UTC),
			},
			expected: []time.Time{},
		},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(fmt.Sprintf("%v", tc.interval), func(t *testing.T) {
			t.Parallel()

			actual := tc.interval.Days()
			assert.Equal(t, tc.expected, actual)
		})
	}
}

func TestInterval_Has(t *testing.T) {
	t.Parallel()

	val, err := time.Parse(time.RFC3339, "2024-01-02T01:02:03Z")
	assert.NoError(t, err)
	log.Println("parsed: ", val)

	truncated := val.UTC().Truncate(time.Hour * 24)
	log.Println("truncated: ", truncated)
}
