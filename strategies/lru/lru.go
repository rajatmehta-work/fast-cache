package lru

import (
	"time"

	"github.com/rajatmehta-work/fast-cache/models"
	"github.com/rajatmehta-work/go-lib/hashmap"
	"github.com/rajatmehta-work/go-lib/list"
)

type lruMem[K comparable, V any] struct {
	seq     *list.List[K, V]
	present hashmap.Map[K, *list.Element[K, V]]
	opts    options
}

func New[K comparable, V any](opts ...Option) models.Memory[K, V] {
	lru := lruMem[K, V]{}
	lru.seq = list.New[K, V]()
	lru.present = hashmap.New[K, *list.Element[K, V]]()

	for _, opt := range opts {
		opt(&lru.opts)
	}
	return lru
}

func (l lruMem[K, V]) Get(key K) (zero V, _ bool) {
	ref, ok := l.present.Get(key)
	if ok {
		l.seq.MoveToFront(ref)
		return ref.Value, true
	}
	return
}
func (l lruMem[K, V]) Set(key K, value V) {
	if ref, ok := l.present.Get(key); !ok {
		if l.opts.capacity == 0 {
			lruRef := l.seq.Back()
			l.present.Delete(lruRef.Key)
			l.seq.Remove(lruRef)
			l.opts.capacity++
		}
		l.opts.capacity--
		l.present.Set(key, l.seq.PushFront(value))
	} else {
		ref.Value = value
		l.seq.MoveToFront(ref)
	}
}
func (l lruMem[K, V]) Delete(key K) {
	if ref, ok := l.present.Get(key); ok {
		l.seq.Remove(ref)
		l.present.Delete(key)
		l.opts.capacity++
	}
}
func (l lruMem[K, V]) SetWithTTL(key K, value V, time time.Duration) {}
