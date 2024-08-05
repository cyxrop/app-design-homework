package order

import (
	orderEntity "applicationDesignTest/internal/entity/order"
	"applicationDesignTest/internal/entity/room"
	"context"
	"fmt"
)

// TODO: add tests
func (s *Service) CreateOrder(ctx context.Context, order orderEntity.Order) error {
	err := order.Validate()
	if err != nil {
		return fmt.Errorf("%w: %w", orderEntity.ErrInvalidOrder, err)
	}

	return s.tm.InTx(ctx, func(ctx context.Context) (err error) {
		availabilities, err := s.roomRepository.GetRoomAvailability(ctx, order.HotelID, order.RoomID)
		if err != nil {
			return fmt.Errorf("get room availability: %w", err)
		}

		unavailable := order.TryReserve(availabilities)
		if len(unavailable) > 0 {
			return fmt.Errorf("%w: dates: %v", room.ErrRoomIsNotAvailable, unavailable)
		}

		err = s.roomRepository.UpdateRoomAvailability(ctx, order.HotelID, order.RoomID, availabilities)
		if err != nil {
			return fmt.Errorf("save room availability: %w", err)
		}

		err = s.orderRepository.CreateOrder(ctx, order)
		if err != nil {
			return fmt.Errorf("create order: %w", err)
		}
		return nil
	})
}
