package cache

import "time"

// err options constants
const (
	ErrInCorrectOption = "provided options are invalid"
)

type Option func(*options)

type options struct {
	capacity  uint32
	evictAlgo int
	ttl       time.Duration
}

func (o options) IsValid() bool {
	for _, evictAl := range CacheEvictionAlgorithms {
		if evictAl == o.evictAlgo {
			return true
		}
	}
	return false
}

func WithCapacity(c uint32) Option {
	return func(o *options) {
		o.capacity = c
	}
}

func WithEvictionAlgo(a int) Option {
	return func(o *options) {
		o.evictAlgo = a
	}
}

func WithTTL(dur time.Duration) Option {
	return func(o *options) {
		o.ttl = dur
	}
}
