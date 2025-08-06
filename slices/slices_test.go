package slices

import "testing"

func TestCompare(t *testing.T) {
	input1 := []int{1, 2, 3}
	input2 := []int{1, 2, 3}
	input3 := []int{4, 5, 6}
	input4 := []int{1, 2, 3, 4}

	if !Compare(input1, input2) {
		t.Errorf("Expected slices to be equal: %v and %v", input1, input2)
	}

	if !Compare(input1, input1) {
		t.Errorf("Expected slices to be equal: %v and itself", input1)
	}

	if Compare(input1, input3) {
		t.Errorf("Expected slices to be different: %v and %v", input1, input3)
	}

	if Compare(input1, input4) {
		t.Errorf("Expected slices to be different: %v and %v", input1, input4)
	}
}

func TestCompareN(t *testing.T) {
	input1 := []int{1, 2, 3, 4}
	input2 := []int{1, 2, 3, 4, 5}
	input3 := []int{1, 2, 4}

	if !CompareN(input1, input2, 4) {
		t.Errorf("Expected first %d elements to be equal: %v and %v", 4, input1, input2)
	}

	if !CompareN(input1, input3, 2) {
		t.Errorf("Expected first %d elements to be different: %v and %v", 2, input1, input3)
	}

	if CompareN(input1, input3, 3) {
		t.Errorf("Expected first %d elements to be different: %v and %v", 3, input1, input3)
	}

	if CompareN(input1, input2, 5) {
		t.Errorf("Expected comparison with n greater than slice length to return false")
	}
}
