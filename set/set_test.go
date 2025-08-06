package set

import (
	"testing"
)

func TestSetInsert(t *testing.T) {
	s := New[string]()
	if s.Len() != 0 {
		t.Errorf("Expected set to be empty, got length %d", s.Len())
	}

	s.Insert("foo")
	if s.Len() != 1 {
		t.Errorf("Expected set to have length 1, got %d", s.Len())
	}

	s.Insert("bar")
	if s.Len() != 2 {
		t.Errorf("Expected set to have length 2, got %d", s.Len())
	}
}

func TestSetNewWithValues(t *testing.T) {
	s := NewWithValues("foo", "bar", "baz")
	if s.Len() != 3 {
		t.Errorf("Expected set to have length 3, got %d", s.Len())
	}

	if !s.Contains("foo") || !s.Contains("bar") || !s.Contains("baz") {
		t.Error("Expected set to contain 'foo', 'bar', and 'baz'")
	}

	if s.Contains("qux") {
		t.Error("Expected set to not contain 'qux'")
	}
}

func TestSetContains(t *testing.T) {
	s := New[string]()
	if s.Contains("foo") {
		t.Errorf("Expected set to not contain 'foo'")
	}

	s.Insert("foo")
	if !s.Contains("foo") {
		t.Errorf("Expected set to contain 'foo'")
	}

	if s.Contains("bar") {
		t.Errorf("Expected set to not contain 'bar'")
	}
}

func TestSetIsEmpty(t *testing.T) {
	s := New[string]()
	if !s.IsEmpty() {
		t.Error("Expected new set to be empty")
	}

	s.Insert("foo")
	if s.IsEmpty() {
		t.Error("Expected set to not be empty after inserting an element")
	}

	s.Clear()
	if !s.IsEmpty() {
		t.Error("Expected set to be empty after clearing")
	}
}

func TestSetRemove(t *testing.T) {
	s := New[string]()
	s.Insert("foo")
	s.Insert("bar")

	s.Remove("foo")
	if s.Contains("foo") {
		t.Errorf("Expected set to not contain 'foo' after removal")
	}

	if !s.Contains("bar") {
		t.Errorf("Expected set to still contain 'bar'")
	}

	if s.Len() != 1 {
		t.Errorf("Expected set to have length 1 after removing 'foo', got %d", s.Len())
	}

	s.Remove("foobar")
	if !s.Contains("bar") {
		t.Errorf("Expected set to still contain 'bar'")
	}

	if s.Len() != 1 {
		t.Errorf("Expected set length to remain 1 after removing non-existent element, got %d", s.Len())
	}
}

func TestSetClear(t *testing.T) {
	s := NewWithValues("foo", "bar", "baz")
	if s.Len() != 3 {
		t.Errorf("Expected set to have length 3, got %d", s.Len())
	}

	s.Clear()
	if !s.IsEmpty() {
		t.Error("Expected set to be empty after clearing")
	}

	if s.Len() != 0 {
		t.Errorf("Expected set length to be 0 after clearing, got %d", s.Len())
	}
}

func TestSetAll(t *testing.T) {
	s1 := NewWithValues("foo", "bar", "baz")
	s2 := New[string]()

	for elem := range s1.All() {
		s2.Insert(elem)
	}

	if s2.Len() != 3 {
		t.Errorf("Expected s2 to have length 3 after inserting all elements from s, got %d", s2.Len())
	}

	for _, elem := range []string{"foo", "bar", "baz"} {
		if !s2.Contains(elem) {
			t.Errorf("Expected s2 to contain '%s'", elem)
		}
	}
}
