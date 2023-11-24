package simple

import "time"

type Option func(*options)

type options struct {
	ttl time.Duration
}

func WithTTL(d time.Duration) Option {
	return func(o *options) {
		o.ttl = d
	}
}
