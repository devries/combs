package combs

// Creates a channel that returns all combinations of n items from v
func Combinations[T any](n int, v []T) <-chan []T {
	ch := make(chan []T)
	go func() {
		prefix := []T{}
		stringCombinationsRecursor(prefix, n, v, ch)
		close(ch)
	}()

	return ch
}

func stringCombinationsRecursor[T any](prefix []T, n int, v []T, ch chan<- []T) {
	if n == 1 {
		for _, c := range v {
			l := len(prefix)
			result := make([]T, l+1)
			copy(result, prefix)
			result[l] = c
			ch <- result
		}
		return
	}

	if n == len(v) {
		result := make([]T, len(prefix)+len(v))
		copy(result, prefix)
		copy(result[len(prefix):], v)
		ch <- result
		return
	}

	l := len(prefix)
	newPrefix := make([]T, l+1)

	copy(newPrefix, prefix)
	newPrefix[l] = v[0]

	// combinations with this element
	stringCombinationsRecursor(newPrefix, n-1, v[1:], ch)

	// combinations without this element
	stringCombinationsRecursor(prefix, n, v[1:], ch)
}
