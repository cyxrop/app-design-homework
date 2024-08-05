package room

import (
	"applicationDesignTest/internal/entity/room"
	"sync"
	"time"
)

type roomAvailability struct {
	mx             *sync.Mutex
	availabilities room.Availabilities
}

type Repository struct {
	availabilities map[string]map[string]*roomAvailability
}

func NewRepository() *Repository {
	return &Repository{
		availabilities: map[string]map[string]*roomAvailability{
			"reddison": {
				"lux": {
					mx: &sync.Mutex{},
					availabilities: room.Availabilities{
						{"reddison", "lux", time.Date(2024, 1, 1, 0, 0, 0, 0, time.UTC), 1},
						{"reddison", "lux", time.Date(2024, 1, 2, 0, 0, 0, 0, time.UTC), 1},
						{"reddison", "lux", time.Date(2024, 1, 3, 0, 0, 0, 0, time.UTC), 1},
						{"reddison", "lux", time.Date(2024, 1, 4, 0, 0, 0, 0, time.UTC), 1},
						{"reddison", "lux", time.Date(2024, 1, 5, 0, 0, 0, 0, time.UTC), 0},
					},
				},
			},
		},
	}
}
