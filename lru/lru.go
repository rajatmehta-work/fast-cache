package lru

import "sync"

type LRU[K comparable, V any] interface {
	Get(key K) V
	Put(key K, value V)
	Del(key K)
	PutTTL(key K, value V, time int) //time in sec
}

type lryMem[K comparable, V any] struct {
	options options
	seq Sequence[V any]
	present Present[K comparable]
	mut     sync.Mutex
}

func New(opt ...Options) LRU {
	lru := lryMem{}
	applyOptions(lru)
	return LRU(lru)
}

func (l lryMem) Get(key K) V {
	ref,ok:= l.present.GetValue[K](key)
	if ok{
		val:=l.GetValue[V](ref)
		l.Evict[V](ref)
		l.Append[V](ref)
		return val
	}
	return nil
}
func (l lryMem) Put(key K, value V) {
	newRef:=l.seq.NewRef()
	ref,ok:= l.present.GetValue[K](key)
	if ok{
		l.Evict[V](ref)
		l.APpend(newRef)
		l.present.Update(key,newRef)
	}else{
		if l.capacity==len(l.present.Len()){
			l.sequence.FirstDelete()
			l.sequence.Append()
		}else{
			l.put(key,newRef)
			l.sequence.Append(newRef)
		}
	}
	
5

	return nil
}
func (l lryMem) Del(key K) {

}
func (l lryMem) PutTTL(key K, value V, time int) {

}
