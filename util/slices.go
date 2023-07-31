package util

func ReduceArray[T, M any](s []T, f func(M, T) M, initValue M) M {
	acc := initValue
	for _, v := range s {
		acc = f(acc, v)
	}
	return acc
}

func FilterArray[T any](d, s []T, keep func(i T) bool) []T {
	for _, n := range s {
		if keep(n) {
			d = append(d, n)
		}
	}
	return d
}
