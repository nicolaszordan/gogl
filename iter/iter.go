package iter

import "iter"

func Filter[V any](pred func(V) bool, it iter.Seq[V]) iter.Seq[V] {
	return func(yield func(V) bool) {
		for elem := range it {
			if pred(elem) {
				if !yield(elem) {
					return
				}
			}
		}
	}
}

// Unzip takes a sequence of pairs and splits it into two sequences.
func Unzip[K, V any](it iter.Seq2[K, V]) (iter.Seq[K], iter.Seq[V]) {
	return func(yield func(K) bool) {
			for k := range it {
				if !yield(k) {
					return
				}
			}
		}, func(yield func(V) bool) {
			for _, v := range it {
				if !yield(v) {
					return
				}
			}
		}
}
