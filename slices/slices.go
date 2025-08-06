package slices

// Compare checks if two slices are equal.
func Compare[T comparable](s1, s2 []T) bool {
	if len(s1) != len(s2) {
		return false
	}

	for i := range s1 {
		if s1[i] != s2[i] {
			return false
		}
	}
	return true
}

// CompareN compares the first n elements of two slices.
func CompareN[T comparable](s1, s2 []T, n int) bool {
	if n > len(s1) || n > len(s2) {
		return false
	}

	for i := range n {
		if s1[i] != s2[i] {
			return false
		}
	}
	return true
}
