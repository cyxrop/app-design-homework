package order

import (
	"applicationDesignTest/internal/entity/order"
	"sync"
)

type Repository struct {
	mx     *sync.Mutex
	orders []order.Order
}

func NewRepository() *Repository {
	return &Repository{
		mx:     &sync.Mutex{},
		orders: []order.Order{},
	}
}
