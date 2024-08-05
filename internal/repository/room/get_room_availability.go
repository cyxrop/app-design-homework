package room

import (
	roomEntity "applicationDesignTest/internal/entity/room"
	"applicationDesignTest/internal/tx/lock_storage"
	"context"
)

func (r *Repository) GetRoomAvailability(ctx context.Context, hotelId, roomId string) (roomEntity.Availabilities, error) {
	ls, err := lock_storage.FromContext(ctx)
	if err != nil {
		return nil, err
	}

	rooms, ok := r.availabilities[hotelId]
	if !ok {
		return nil, roomEntity.ErrHotelIsNotFound
	}
	val, ok := rooms[roomId]
	if !ok {
		return nil, roomEntity.ErrRoomIsNotFound
	}

	ls.LockAndStore(val.mx)

	list := make(roomEntity.Availabilities, len(val.availabilities))
	copy(list, val.availabilities)
	return list, nil
}
