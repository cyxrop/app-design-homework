package order

import (
	"applicationDesignTest/internal/entity/date"
	orderEntity "applicationDesignTest/internal/entity/order"
	"applicationDesignTest/internal/entity/room"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

func (s *HttpService) CreateOrder(w http.ResponseWriter, r *http.Request) {
	var (
		ctx = r.Context()
		req createOrder
	)
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		s.writeError(w, fmt.Errorf("failed to decode request: %s", err), http.StatusBadRequest)
		return
	}

	err = req.Validate()
	if err != nil {
		s.writeError(w, fmt.Errorf("invalid request: %s", err), http.StatusBadRequest)
		return
	}

	newOrder := req.ToOrder()
	err = s.orderSvc.CreateOrder(ctx, newOrder)
	if err != nil {
		status := http.StatusInternalServerError
		switch {
		case errors.Is(err, orderEntity.ErrInvalidOrder),
			errors.Is(err, room.ErrHotelIsNotFound),
			errors.Is(err, room.ErrRoomIsNotFound),
			errors.Is(err, room.ErrRoomIsNotAvailable):
			status = http.StatusBadRequest
			s.writeError(w, fmt.Errorf("failed to create order: %s", err), status)
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err = json.NewEncoder(w).Encode(newOrder); err != nil {
		http.Error(w, fmt.Sprintf("failed to encode response: %s", err), http.StatusInternalServerError)
		return
	}

	s.logger.Info(fmt.Sprintf("Order successfully created: %v", newOrder))
}

type createOrder struct {
	HotelID   string    `json:"hotel_id"`
	RoomID    string    `json:"room_id"`
	UserEmail string    `json:"email"`
	From      time.Time `json:"from"`
	To        time.Time `json:"to"`
}

func (o createOrder) Validate() error {
	if o.HotelID == "" {
		return errors.New("hotel id is empty")
	}
	if o.RoomID == "" {
		return errors.New("room id is empty")
	}
	if o.UserEmail == "" {
		return errors.New("user email is empty")
	}
	if o.From.IsZero() {
		return errors.New("from is empty")
	}
	if o.To.IsZero() {
		return errors.New("to is empty")
	}
	if o.From.After(o.To) {
		return errors.New("from is after to")
	}
	return nil
}

func (o createOrder) ToOrder() orderEntity.Order {
	return orderEntity.Order{
		HotelID:   o.HotelID,
		RoomID:    o.RoomID,
		UserEmail: o.UserEmail,
		Interval: date.Interval{
			From: o.From.UTC().Truncate(24 * time.Hour),
			To:   o.To.UTC().Truncate(24 * time.Hour),
		},
	}
}
