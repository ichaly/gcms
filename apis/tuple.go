package apis

type Tuple[T comparable, K comparable, V any] struct {
	keys   map[T]K
	values map[K]V
}

func (my Tuple[T, K, V]) GetByK(k any) (V, bool) {
	val, ok := my.keys[k]
	if !ok {
		return (V(nil)), false
	}
	val, ok := my.values[val]
	return val, ok
}
