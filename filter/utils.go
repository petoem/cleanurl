package filter

func toSet[E comparable](slice []E) map[E]struct{} {
	set := make(map[E]struct{})
	for _, v := range slice {
		set[v] = struct{}{}
	}
	return set
}

func mapFunc[E any](slice []E, f func(E) E) {
	for i, v := range slice {
		slice[i] = f(v)
	}
}
