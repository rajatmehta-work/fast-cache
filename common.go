package fastcache

type Memory interface {
	Get()
	Set()
	Delete()
	SetWithTTL()
}
