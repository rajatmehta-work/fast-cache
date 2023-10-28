package doublyll

type DDLNode[V any] *node[V]

type node[V any] struct {
	prev, next DDLNode[V]
	value      V
}

func (d node[V]) Value() V {
	return d.value
}
