package room

import (
	roomEntity "applicationDesignTest/internal/entity/room"
	"applicationDesignTest/internal/tx/lock_storage"
	"context"
)

func (r *Repository) UpdateRoomAvailability(ctx context.Context, hotelId, roomId string, availabilities roomEntity.Availabilities) error {
	ls, err := lock_storage.FromContext(ctx)
	if err != nil {
		return err
	}

	rooms, ok := r.availabilities[hotelId]
	if !ok {
		return roomEntity.ErrHotelIsNotFound
	}
	val, ok := rooms[roomId]
	if !ok {
		return roomEntity.ErrRoomIsNotFound
	}

	ls.LockAndStore(val.mx)
	val.availabilities = availabilities
	return nil
}
