package cache

import (
	"errors"
	"time"

	fastcache "github.com/rajatmehta-work/fast-cache"
)

type FastCache interface {
	Get()
	Set()
	SetWithTTL(time.Duration)
	Delete()
}

type cache struct {
	data    fastcache.Memory
	options *options
}

func New(opts ...Option) (FastCache, error) {
	c := cache{}

	for _, opt := range opts {
		opt(c.options)
	}
	if !c.options.IsValid() {
		return FastCache(c), errors.New(ErrInCorrectOption)
	}
	return FastCache(c), nil
}

func (fc cache) Get() {

	if fc.options.evictAlgo == LRU {

	} else if fc.options.evictAlgo == LFU {

	} else {

	}
}
func (fc cache) Set() {}

func (fc cache) SetWithTTL(dur time.Duration) {}

func (fc cache) Delete() {}
