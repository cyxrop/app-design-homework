package order

import (
	"applicationDesignTest/internal/entity/order"
	"applicationDesignTest/internal/tx/lock_storage"
	"context"
)

func (r *Repository) CreateOrder(ctx context.Context, order order.Order) error {
	ls, err := lock_storage.FromContext(ctx)
	if err != nil {
		return err
	}

	ls.LockAndStore(r.mx)
	r.orders = append(r.orders, order)
	return nil
}
