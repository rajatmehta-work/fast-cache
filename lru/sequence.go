package lru

type Sequence[V any] interface {
	Append(value V) *ddNode
	Evict(node *ddNode)
	GetValue(node *ddNode)
}

type dll struct {
	dLLHead *ddNode[V]
	dLLTail *ddNode[V]
}


func(s *dll)Append[V any](value V)*ddNode{
	newNode:=&ddNode[V]{value:value}
	if dLLHead==nil{
		dLLHead=newNode
		dLLTail=newNode
	}else{
		newNode.prev=dLLTail
		dLLTail=newNode
	}
}

func(s *dll)Evict(node *ddNode){
	if dLLHead==node{
		dLLHead=nil
		dLLTail=nil
	}else if dLLTail==node{
		dLLTail=node.Prev
		node.Prev.Next=nil
	}else{
		node.Prev.Next=node.Next
		node.next.prev=node.prev
	}
		// delete(node)
}

func(s *dll)GetValue[V any](node *ddNode)(V,error){
	if node!=nil{
		return nil,fmt.Erorrs("empty node")
	}
	return node.value,nil

}