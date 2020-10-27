package calculate

import (
	"fmt"
	"testing"

	calculate "github.com/ericcchiu/golang-intro/go-testing/calculate"
)

func TestAdd(t *testing.T) {
	// testing.Short is a flag that tells Go to skip the current test without having to comment out test.
	// if testing.Short() {
	// 	t.SkipNow()
	// }

	addTests := []struct {
		a        int
		b        int
		expected int
	}{
		{a: 0, b: 0, expected: 0},
		{a: 0, b: -5, expected: -5},
		{a: 150, b: 55, expected: 205},
	}

	for i, test := range addTests {
		name := fmt.Sprintf("Test #%v", i)
		t.Run(name, func(t *testing.T) {
			result := calculate.Add(test.a, test.b)
			if result != test.expected {
				t.Errorf("Expected %v, but got %v", test.expected, result)
			}
		})
		// result := calculate.Add(test.a, test.b)
		// if result != test.expected {
		// 	t.Errorf("Expected %v, but got %v", test.expected, result)
		// }
	}

	result := calculate.Add(3, 2)
	if result != 5 {
		t.Errorf("Expected 5, got %v", result)
	}
}
func TestSubtract(t *testing.T) {
	result := calculate.Subtract(3, 2)
	if result != 1 {
		t.Errorf("Expected 1, got %v", result)
	}
}
