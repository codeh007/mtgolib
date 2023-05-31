package utils

func Filter[T any](slice []T, f func(T) bool) []T {
	var n []T
	for _, e := range slice {
		if f(e) {
			n = append(n, e)
		}
	}
	return n
}

func Map[T any](slice []*T, f func(*T) *T) []*T {
	var n []*T
	for _, e := range slice {
		newE := f(e)
		n = append(n, newE)
	}
	return n
}
