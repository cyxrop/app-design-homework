package tx

import (
	"applicationDesignTest/internal/tx/lock_storage"
	"context"
	"sync"
)

type InMemoryTxManager struct {
	locks []*sync.Mutex
}

func NewInMemoryTxManager() *InMemoryTxManager {
	return &InMemoryTxManager{
		locks: make([]*sync.Mutex, 0),
	}
}

func (tm *InMemoryTxManager) InTx(ctx context.Context, f func(ctx context.Context) error) error {
	ls := lock_storage.NewLockStorage()
	defer func() {
		ls.UnlockAll()
	}()

	ctx = lock_storage.ContextWithLockStorage(ctx, ls)
	err := f(ctx)
	return err
}
