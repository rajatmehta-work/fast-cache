package lru

import (
	"time"

	"github.com/rajatmehta-work/fast-cache/models"
	dll "github.com/rajatmehta-work/fast-cache/models/dll"
	storage "github.com/rajatmehta-work/fast-cache/models/storage"
)

type lruMem[K comparable, V any] struct {
	seq     dll.DoublyLL[V]
	present storage.Storage[K, dll.DDLNode[V]]
	opts    options
}

func New[K comparable, V any](opts ...Option) models.Memory[K, V] {
	lru := lruMem[K, V]{}
	lru.seq = dll.New[V]()
	lru.present = storage.New[K, dll.DDLNode[V]]()

	for _, opt := range opts {
		opt(&lru.opts)
	}
	return lru
}

func (l lruMem[K, V]) Get(key K) (V, error) {
	var val V
	var err error
	ref, ok := l.present.Get(key)
	if ok {
		val, err = l.seq.GetValue(ref)
		if err != nil {
			return val, err
		}
		l.seq.Delete(ref)
		// l.seq.Append(ref)
		return val, err
	} else {
		return val, err
	}

}
func (l lruMem[K, V]) Set(key K, value V) {
	_, ok := l.present.Get(key)
	if !ok {

		if l.opts.capacity == 0 {
			l.seq.Delete(l.seq.GetTail())
			// l.present
		} else {
			l.opts.capacity--
			l.seq.Append(value)
			l.present.Set(key, l.seq.GetTail())
			l.seq.Append(value)
		}

	}
}
func (l lruMem[K, V]) Delete(key K) {

}
func (l lruMem[K, V]) SetWithTTL(key K, value V, time time.Duration) {}
