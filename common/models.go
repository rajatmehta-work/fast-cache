package common

import "time"

type Memory[K comparable, V any] interface {
	Get(K) (V, bool)
	Set(K, V)
	Delete(K)
	Len() int
	SetWithTTL(K, V, time.Duration)
}
