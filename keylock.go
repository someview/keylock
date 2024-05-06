package keylock

import (
	"context"
	"sync"

	"golang.org/x/sync/semaphore"
)

type KeyLock struct {
	sync.Map
}

func (l *KeyLock) TryAcquire(key int64) bool {
	val, _ := l.Map.LoadOrStore(key, semaphore.NewWeighted(1))
	return val.(*semaphore.Weighted).TryAcquire(key)
}

func (l *KeyLock) Acquire(ctx context.Context, key int64) error {
	val, _ := l.Map.LoadOrStore(key, semaphore.NewWeighted(1))
	return val.(*semaphore.Weighted).Acquire(ctx, 1)
}

func (l *KeyLock) Release(key int64) {
	val, _ := l.Map.LoadAndDelete(key)
	val.(*semaphore.Weighted).Release(1)
}
