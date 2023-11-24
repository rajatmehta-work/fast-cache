package simple

import (
	"time"

	"github.com/rajatmehta-work/fast-cache/common"
	"github.com/rajatmehta-work/go-lib/hashmap"
)

type simpleMem[K comparable, V any] struct {
	present hashmap.Map[K, V]
	opts    options
}

func New[K comparable, V any](opts ...Option) common.Memory[K, V] {
	lru := simpleMem[K, V]{}
	lru.present = hashmap.New[K, V]()
	for _, opt := range opts {
		opt(&lru.opts)
	}
	return &lru
}

func (l *simpleMem[K, V]) Get(key K) (V, bool) {
	val, ok := l.present.Get(key)
	return val, ok
}
func (l *simpleMem[K, V]) Set(key K, value V) {
	l.present.Set(key, value)
}

func (l *simpleMem[K, V]) Delete(key K) {
	if _, ok := l.present.Get(key); ok {
		l.present.Delete(key)
	}
}

func (l *simpleMem[K, V]) Len() int {
	return l.present.Len()
}

func (l *simpleMem[K, V]) SetWithTTL(key K, value V, time time.Duration) {}
