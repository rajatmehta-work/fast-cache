package doublyll

type DoublyLL[V any] interface {
	GetHead() DDLNode[V]
	GetTail() DDLNode[V]
	Append(V)
	Delete(DDLNode[V])
	GetValue(DDLNode[V]) (V, error)
}

type dll[V any] struct {
	dLLHead DDLNode[V]
	dLLTail DDLNode[V]
}

func New[V any]() DoublyLL[V] {
	return dll[V]{}
}

func (s dll[V]) Append(v V) {
	newNode := &node[V]{value: v}
	if s.dLLHead == nil {
		s.dLLHead = newNode
	} else {
		newNode.prev = s.dLLTail
	}
	s.dLLTail = newNode
}

func (s dll[V]) Delete(node DDLNode[V]) {
	if s.dLLHead == node {
		if s.dLLTail == node {
			s.dLLTail = nil
		}
		s.dLLHead = s.dLLHead.next
	} else if s.dLLTail == node {
		s.dLLTail = s.dLLTail.prev
		s.dLLTail.next = nil
	} else {
		node.prev.next = node.next
		node.next.prev = node.prev
	}
	node = nil // freeing memory will be tc by GC
}

func (s dll[V]) GetValue(node DDLNode[V]) (V, error) {
	return node.value, nil
}

func (s dll[V]) GetHead() DDLNode[V] {
	return s.dLLHead
}

func (s dll[V]) GetTail() DDLNode[V] {
	return s.dLLTail
}
