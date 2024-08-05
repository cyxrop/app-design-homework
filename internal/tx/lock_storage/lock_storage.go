package lock_storage

import "sync"

type LockStorage struct {
	mx       *sync.RWMutex
	locks    []*sync.Mutex
	locksMap map[*sync.Mutex]struct{}
}

func NewLockStorage() *LockStorage {
	return &LockStorage{
		mx:       &sync.RWMutex{},
		locks:    make([]*sync.Mutex, 0),
		locksMap: make(map[*sync.Mutex]struct{}),
	}
}

func (ls *LockStorage) LockAndStore(lock *sync.Mutex) {
	ls.mx.Lock()
	defer ls.mx.Unlock()

	_, ok := ls.locksMap[lock]
	if ok {
		return
	}

	lock.Lock()
	ls.locks = append(ls.locks, lock)
	ls.locksMap[lock] = struct{}{}
}

func (ls *LockStorage) UnlockAll() {
	ls.mx.Lock()
	defer ls.mx.Unlock()
	for _, l := range ls.locks {
		l.Unlock()
	}
}
