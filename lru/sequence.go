package lru

type sequence[V any] interface {
	append(V)
	delete(ddlNode[V])
	getValue(ddlNode[V]) (V, error)
}

type dll[V any] struct {
	dLLHead ddlNode[V]
	dLLTail ddlNode[V]
}

func newSeq[V any]() sequence[V] {
	return dll[V]{}
}

func (s dll[V]) append(v V) {
	newNode := &node[V]{value: v}
	if s.dLLHead == nil {
		s.dLLHead = newNode
		s.dLLTail = newNode
	} else {
		newNode.prev = s.dLLTail
		s.dLLTail = newNode
	}
}

func (s dll[V]) delete(node ddlNode[V]) {
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

func (s dll[V]) getValue(node ddlNode[V]) (V, error) {
	return node.value, nil
}
