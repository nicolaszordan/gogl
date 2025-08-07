package iter

import (
	"slices"
	"testing"

	glslices "github.com/nicolaszordan/gogl/slices"
)

func TestFilter(t *testing.T) {
	input_slice := []int{1, 2, 3, 4, 5}
	is_even := func(n int) bool {
		return n%2 == 0
	}

	expected := []int{2, 4}

	output := make([]int, 0)
	for v := range Filter(is_even, slices.Values(input_slice)) {
		output = append(output, v)
	}

	if !glslices.Compare(expected, output) {
		t.Errorf("Expected %v but got %v", expected, output)
	}
}

func TestUnzip(t *testing.T) {
	input := []string{"one", "two", "three"}

	outputItA, outputItB := Unzip(slices.All(input))

	expectedA := []int{0, 1, 2}
	expectedB := []string{"one", "two", "three"}

	var outputA []int
	var outputB []string
	for v := range outputItB {
		outputB = append(outputB, v)
	}
	for v := range outputItA {
		outputA = append(outputA, v)
	}
	if !glslices.Compare(expectedA, outputA) {
		t.Errorf("Expected %v but got %v for first output", expectedA, outputA)
	}
	if !glslices.Compare(expectedB, outputB) {
		t.Errorf("Expected %v but got %v for second output", expectedB, outputB)
	}
}
