package arrays

import (
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	input := []int{1, 1, 2, 2, 2, 3, 4}
	expected := []int{1, 2, 3, 4}

	result := arrays.RemoveDuplicates(input)

	if result != expected {
		t.Errorf("%t != %t", result, expected)
	}
}
