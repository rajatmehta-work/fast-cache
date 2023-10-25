package lru

type Option interface {
	Apply(*options)
}

type optionFunc func(*options)

type options struct {
	Capacity int
	///
}

func applyOptions(parOpt options, options ...Option) {
	for _, opt := range options {
		opt(parOpt)
	}
}
func (fn Option) Apply() {

}

func WithCapacity(cap int) Options {
	return optionsFunc(func(o *options) {
		o.Capacity = cap
	})
}
