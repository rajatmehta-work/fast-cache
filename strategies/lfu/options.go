package lfu

import "time"

type Option func(*options)

type options struct {
	capacity int
	ttl      time.Duration
	asc      bool
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

func WithAscOrderEvictionPolicy() Option {
	return func(o *options) {
		o.asc = true
	}
}
