package set

import "iter"

type Set[T comparable] struct {
	set map[T]struct{}
}

func New[T comparable]() *Set[T] {
	return &Set[T]{
		set: make(map[T]struct{}),
	}
}

func NewWithValues[T comparable](elems ...T) *Set[T] {
	set := New[T]()
	for _, elem := range elems {
		set.Insert(elem)
	}
	return set
}

func (s *Set[T]) Len() int {
	return len(s.set)
}

func (s *Set[T]) Insert(elem T) {
	s.set[elem] = struct{}{}
}

func (s *Set[T]) Remove(elem T) {
	delete(s.set, elem)
}

func (s *Set[T]) IsEmpty() bool {
	return s.Len() == 0
}

func (s *Set[T]) Contains(needle T) bool {
	_, ok := s.set[needle]
	return ok
}

func (s *Set[T]) Clear() {
	s.set = make(map[T]struct{})
}

func (s *Set[T]) All() iter.Seq[T] {
	return func(yield func(T) bool) {
		for v := range s.set {
			if !yield(v) {
				return
			}
		}
	}
}
