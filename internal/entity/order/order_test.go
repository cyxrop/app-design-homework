package order

import (
	"applicationDesignTest/internal/entity/date"
	"applicationDesignTest/internal/entity/room"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestOrder_Validate(t *testing.T) {
	t.Parallel()

	t.Run("positive", func(t *testing.T) {
		t.Parallel()

		order := Order{
			HotelID:   "hotel-1",
			RoomID:    "room-1",
			UserEmail: "user-1",
			Interval: date.Interval{
				From: time.Date(2024, 8, 1, 0, 0, 0, 0, time.UTC),
				To:   time.Date(2024, 8, 3, 0, 0, 0, 0, time.UTC),
			},
		}

		err := order.Validate()
		assert.NoError(t, err)
	})

	t.Run("negative", func(t *testing.T) {
		t.Parallel()

		cases := []struct {
			name     string
			order    Order
			expected error
		}{
			{
				name: "invalid hotel id",
				order: Order{
					HotelID:   "",
					RoomID:    "room-1",
					UserEmail: "user-1",
					Interval: date.Interval{
						From: time.Date(2024, 8, 1, 0, 0, 0, 0, time.UTC),
						To:   time.Date(2024, 8, 3, 0, 0, 0, 0, time.UTC),
					},
				},
				expected: errors.New("empty hotel id"),
			},
			{
				name: "invalid room id",
				order: Order{
					HotelID:   "hotel-1",
					RoomID:    "",
					UserEmail: "user-1",
					Interval: date.Interval{
						From: time.Date(2024, 8, 1, 0, 0, 0, 0, time.UTC),
						To:   time.Date(2024, 8, 3, 0, 0, 0, 0, time.UTC),
					},
				},
				expected: errors.New("empty room id"),
			},
			{
				name: "invalid user email",
				order: Order{
					HotelID:   "hotel-1",
					RoomID:    "room-1",
					UserEmail: "",
					Interval: date.Interval{
						From: time.Date(2024, 8, 1, 0, 0, 0, 0, time.UTC),
						To:   time.Date(2024, 8, 3, 0, 0, 0, 0, time.UTC),
					},
				},
				expected: errors.New("empty user email"),
			},
			{
				name: "from is empty",
				order: Order{
					HotelID:   "hotel-1",
					RoomID:    "room-1",
					UserEmail: "user-1",
					Interval: date.Interval{
						From: time.Time{},
						To:   time.Date(2024, 8, 1, 0, 0, 0, 0, time.UTC),
					},
				},
				expected: errors.New("from is empty"),
			},
			{
				name: "to is empty",
				order: Order{
					HotelID:   "hotel-1",
					RoomID:    "room-1",
					UserEmail: "user-1",
					Interval: date.Interval{
						From: time.Date(2024, 8, 1, 0, 0, 0, 0, time.UTC),
						To:   time.Time{},
					},
				},
				expected: errors.New("to is empty"),
			},
			{
				name: "from is after to",
				order: Order{
					HotelID:   "hotel-1",
					RoomID:    "room-1",
					UserEmail: "user-1",
					Interval: date.Interval{
						From: time.Date(2024, 8, 3, 0, 0, 0, 0, time.UTC),
						To:   time.Date(2024, 8, 1, 0, 0, 0, 0, time.UTC),
					},
				},
				expected: errors.New("from is after to"),
			},
		}

		for _, tc := range cases {
			tc := tc
			t.Run(tc.name, func(t *testing.T) {
				t.Parallel()

				actual := tc.order.Validate()
				assert.Equal(t, tc.expected, actual)
			})
		}
	})
}

func TestOrder_TryReserve(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name                   string
		order                  Order
		availabilities         room.Availabilities
		expected               []time.Time
		expectedAvailabilities room.Availabilities
	}{
		{
			name: "found unavailable days",
			order: Order{
				HotelID:   "hotel-1",
				RoomID:    "room-1",
				UserEmail: "user-1",
				Interval: date.Interval{
					From: time.Date(2024, 8, 1, 0, 0, 0, 0, time.UTC),
					To:   time.Date(2024, 8, 3, 0, 0, 0, 0, time.UTC),
				},
			},
			availabilities: room.Availabilities{
				{
					HotelID: "hotel-1",
					RoomID:  "room-1",
					Date:    time.Date(2024, 8, 1, 0, 0, 0, 0, time.UTC),
					Quota:   2,
				},
				{
					HotelID: "hotel-1",
					RoomID:  "room-1",
					Date:    time.Date(2024, 8, 2, 0, 0, 0, 0, time.UTC),
					Quota:   0,
				},
				{
					HotelID: "hotel-1",
					RoomID:  "room-1",
					Date:    time.Date(2024, 8, 3, 0, 0, 0, 0, time.UTC),
					Quota:   1,
				},
			},
			expected: []time.Time{
				time.Date(2024, 8, 2, 0, 0, 0, 0, time.UTC),
			},
			expectedAvailabilities: room.Availabilities{
				{
					HotelID: "hotel-1",
					RoomID:  "room-1",
					Date:    time.Date(2024, 8, 1, 0, 0, 0, 0, time.UTC),
					Quota:   1,
				},
				{
					HotelID: "hotel-1",
					RoomID:  "room-1",
					Date:    time.Date(2024, 8, 2, 0, 0, 0, 0, time.UTC),
					Quota:   0,
				},
				{
					HotelID: "hotel-1",
					RoomID:  "room-1",
					Date:    time.Date(2024, 8, 3, 0, 0, 0, 0, time.UTC),
					Quota:   0,
				},
			},
		},
		{
			name: "successful reserved",
			order: Order{
				HotelID:   "hotel-1",
				RoomID:    "room-1",
				UserEmail: "user-1",
				Interval: date.Interval{
					From: time.Date(2024, 8, 1, 0, 0, 0, 0, time.UTC),
					To:   time.Date(2024, 8, 3, 0, 0, 0, 0, time.UTC),
				},
			},
			availabilities: room.Availabilities{
				{
					HotelID: "hotel-1",
					RoomID:  "room-1",
					Date:    time.Date(2024, 8, 1, 0, 0, 0, 0, time.UTC),
					Quota:   2,
				},
				{
					HotelID: "hotel-1",
					RoomID:  "room-1",
					Date:    time.Date(2024, 8, 2, 0, 0, 0, 0, time.UTC),
					Quota:   1,
				},
				{
					HotelID: "hotel-1",
					RoomID:  "room-1",
					Date:    time.Date(2024, 8, 3, 0, 0, 0, 0, time.UTC),
					Quota:   3,
				},
			},
			expected: []time.Time{},
			expectedAvailabilities: room.Availabilities{
				{
					HotelID: "hotel-1",
					RoomID:  "room-1",
					Date:    time.Date(2024, 8, 1, 0, 0, 0, 0, time.UTC),
					Quota:   1,
				},
				{
					HotelID: "hotel-1",
					RoomID:  "room-1",
					Date:    time.Date(2024, 8, 2, 0, 0, 0, 0, time.UTC),
					Quota:   0,
				},
				{
					HotelID: "hotel-1",
					RoomID:  "room-1",
					Date:    time.Date(2024, 8, 3, 0, 0, 0, 0, time.UTC),
					Quota:   2,
				},
			},
		},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			actual := tc.order.TryReserve(tc.availabilities)
			assert.Equal(t, tc.expected, actual)
			assert.Equal(t, tc.expectedAvailabilities, tc.availabilities)
		})
	}
}
