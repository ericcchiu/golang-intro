package arraybasics

import (
	"fmt"
	"testing"

	"github.com/ericcchiu/golang-intro/leetcode/arrays"
	arraybasics "github.com/ericcchiu/golang-intro/leetcode/arrays"
)

func TestRemoveDuplicates(t *testing.T) {
	input := []int{1, 1, 1, 2, 2, 2, 3, 4, 5, 5}
	expected := 5

	result := arraybasics.RemoveDuplicates(input)

	if result != expected {
		t.Errorf("%d != %d", result, expected)
	}
}

func TestMaxProfit(t *testing.T) {
	testParameters := []struct {
		prices   []int
		expected int
	}{
		{prices: []int{7, 1, 5, 3, 6, 4}, expected: 7},
		{prices: []int{1, 2, 3, 4, 5}, expected: 4},
		{prices: []int{1, 2, 3, 4, 5}, expected: 7},
		{prices: []int{7, 6, 4, 3, 1}, expected: 0},
	}

	for _, test := range testParameters {
		testName := fmt.Sprintf("%d", test.prices)
		t.Run(testName, func(t *testing.T) {
			result := arrays.MaxProfit(test.prices)
			if result != test.expected {
				t.Errorf("Expected %d but got %d", test.expected, result)
			}
		})
	}
}
