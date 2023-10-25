package lru

type ddNode[V any] struct {
	prev, next *ddNode
	value      V
}
