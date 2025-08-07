package iter

import (
	"slices"
	"testing"

	glslices "github.com/nicolaszordan/gogl/slices"
)

func TestFilter(t *testing.T) {
	inputSlice := []int{1, 2, 3, 4, 5}
	isEven := func(n int) bool {
		return n%2 == 0
	}
	isOdd := func(n int) bool {
		return n%2 != 0
	}

	expectedEven := []int{2, 4}
	expectedOdd := []int{1, 3, 5}

	outputEven := make([]int, 0, len(expectedEven))
	for v := range Filter(isEven, slices.Values(inputSlice)) {
		outputEven = append(outputEven, v)
	}
	outputOdd := make([]int, 0, len(expectedOdd))
	for v := range Filter(isOdd, slices.Values(inputSlice)) {
		outputOdd = append(outputOdd, v)
	}

	if !glslices.Compare(expectedEven, outputEven) {
		t.Errorf("Expected %v but got %v", expectedEven, outputEven)
	}
	if !glslices.Compare(expectedOdd, outputOdd) {
		t.Errorf("Expected %v but got %v", expectedOdd, outputOdd)
	}
}

func TestUnzip(t *testing.T) {
	input := []string{"one", "two", "three"}

	outputItA, outputItB := Unzip(slices.All(input))

	expectedA := []int{0, 1, 2}
	expectedB := []string{"one", "two", "three"}

	outputA := make([]int, 0, len(expectedA))
	outputB := make([]string, 0, len(expectedB))
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
