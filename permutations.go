package combs

// Permutations of an array of Integers
// Uses Heap's Algorithm (thanks wikipedia)
func Permutations[T any](a []T) <-chan []T {
        ch := make(chan []T)
        go func() {
                k := len(a)
                permutationsRecursor(k, a, ch)
                close(ch)
        }()

        return ch
}

func permutationsRecursor[T any](k int, a []T, ch chan<- []T) {
        if k == 1 {
                output := make([]T, len(a))
                copy(output, a)
                ch <- output
        } else {
                permutationsRecursor(k-1, a, ch)

                for i := 0; i < k-1; i++ {
                        if k%2 == 0 {
                                a[i], a[k-1] = a[k-1], a[i]
                        } else {
                                a[0], a[k-1] = a[k-1], a[0]
                        }
                        permutationsRecursor(k-1, a, ch)
                }
        }
}

