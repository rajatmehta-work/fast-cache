package lru

type Option func(*options)

type options struct {
	capacity int
}

func WithCapacity(cap int) Option {
	return func(o *options) {
		o.capacity = cap
	}
}
