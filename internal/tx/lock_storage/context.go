package lock_storage

import (
	"context"
	"errors"
)

var ErrLockStorageNotFound = errors.New("lock storage not found")

type lockStorageKey struct{}

func ContextWithLockStorage(ctx context.Context, ls *LockStorage) context.Context {
	return context.WithValue(ctx, lockStorageKey{}, ls)
}

func FromContext(ctx context.Context) (*LockStorage, error) {
	ls, ok := ctx.Value(lockStorageKey{}).(*LockStorage)
	if !ok {
		return nil, ErrLockStorageNotFound
	}
	return ls, nil
}
