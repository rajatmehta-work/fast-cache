package lru

type Present[K comparable] interface {
	Delete(K)
	GetValue(K)node *ddNode
}


type present struct {
	ptrRef map[K]*ddNode
}

func (p *present)Delete[K comparable](key K){
	remove(p.ptrRef,key)	
}

func (p *present)GetValue[K comparable](key K)(*ddNode,bool){
	val,ok:= p.ptrRef[key]
	return val,ok
}
func (p *present)Udpate[K comparable](key K)(*ddNode,bool){
	val,ok:= p.ptrRef[key]
	return val,ok
}
