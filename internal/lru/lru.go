package lru

type LRU[K comparable, V any] interface {
	Get(K) (V, bool)
	Set(K, V)
	Len() int
	Delete(K) bool
	Peek(K) (V, bool)
	Flush()
}
