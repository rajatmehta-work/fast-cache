package lru

import "time"

type Option func(*options)

type options struct {
	capacity int
	ttl      time.Duration
}

func WithCapacity(cap int) Option {
	return func(o *options) {
		o.capacity = cap
	}
}

func WithTTL(d time.Duration) Option {
	return func(o *options) {
		o.ttl = d
	}
}
