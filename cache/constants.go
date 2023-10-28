package cache

// cache eviction algorithms
const (
	LRU = iota
	LFU
)

// algo lists
var CacheEvictionAlgorithms []int = []int{LRU, LFU}
