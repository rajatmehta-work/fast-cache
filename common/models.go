package models

import "time"

type Memory[K comparable, V any] interface {
	Get(K) (V, bool)
	Set(K, V)
	Delete(K)
	SetWithTTL(K, V, time.Duration)
}
