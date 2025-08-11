package ptr

// To returns a pointer to the given value.
func To[T any](v T) *T {
	return &v
}

// Equal checks if two pointers are nil or if their dereferenced values are equal.
func Equal[T comparable](a, b *T) bool {
	if a == nil || b == nil {
		return a == b
	}
	return *a == *b
}
