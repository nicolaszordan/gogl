package ptr

import (
	"testing"
)

func TestEqual(t *testing.T) {
	a := To(5)
	b := To(5)
	c := To(10)

	if !Equal(a, b) {
		t.Errorf("Expected %v to equal %v", a, b)
	}

	if Equal(a, c) {
		t.Errorf("Expected %v to not equal %v", a, c)
	}

	if Equal(a, nil) {
		t.Errorf("Expected %v to not equal nil", a)
	}

	a = nil
	if !Equal(a, nil) {
		t.Errorf("Expected nil to equal nil")
	}
}
