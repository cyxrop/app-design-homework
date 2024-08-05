package order

import (
	"applicationDesignTest/internal/entity/order"
	"applicationDesignTest/internal/entity/room"
	"context"
)

type txManager interface {
	InTx(ctx context.Context, f func(ctx context.Context) error) error
}

type orderRepository interface {
	CreateOrder(ctx context.Context, order order.Order) error
}

type roomRepository interface {
	GetRoomAvailability(ctx context.Context, hotelId, roomId string) (room.Availabilities, error)
	UpdateRoomAvailability(ctx context.Context, hotelId, roomId string, availabilities room.Availabilities) error
}

type Service struct {
	tm              txManager
	orderRepository orderRepository
	roomRepository  roomRepository
}

func NewService(tm txManager, orderRepository orderRepository, roomRepository roomRepository) *Service {
	return &Service{
		tm:              tm,
		orderRepository: orderRepository,
		roomRepository:  roomRepository,
	}
}
