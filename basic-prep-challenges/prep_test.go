package challenges

import (
	"fmt"
	"testing"

	"github.com/ericcchiu/golang-intro/basic-prep-challenges/conditionals1"
)

func TestIsOldEnoughToDrink(t *testing.T) {
	testParameters := []struct {
		age      int
		expected bool
	}{
		{age: 20, expected: false},
		{age: 21, expected: true},
		{age: 22, expected: true},
	}

	for _, test := range testParameters {
		name := fmt.Sprintf("%d", test.age)
		t.Run(name, func(t *testing.T) {
			result := conditionals1.IsOldEnoughToDrink(test.age)
			if result != test.expected {
				t.Errorf("%t != %t", result, test.expected)
			}
		})
	}
}

func TestIsOldEnoughToDrive(t *testing.T) {
	testParameters := []struct {
		age      int
		expected bool
	}{
		{age: 15, expected: false},
		{age: 16, expected: true},
		{age: 17, expected: true},
	}

	for _, test := range testParameters {
		name := fmt.Sprintf("%d", test.age)
		t.Run(name, func(t *testing.T) {
			result := conditionals1.IsOldEnoughToDrive(test.age)
			if result != test.expected {
				t.Errorf("%t != %t", result, test.expected)
			}
		})
	}
}

func TestIsOldEnoughToVote(t *testing.T) {
	testParameters := []struct {
		age      int
		expected bool
	}{
		{age: 17, expected: false},
		{age: 18, expected: true},
		{age: 19, expected: true},
	}

	for _, test := range testParameters {
		name := fmt.Sprintf("%d", test.age)
		t.Run(name, func(t *testing.T) {
			result := conditionals1.IsOldEnoughToVote(test.age)
			if result != test.expected {
				t.Errorf("%t != %t", result, test.expected)
			}
		})
	}
}

func TestIsOldEnoughToDrinkAndDrive(t *testing.T) {
	testParameters := []struct {
		age      int
		expected bool
	}{
		{age: 16, expected: false},
		{age: 18, expected: false},
		{age: 21, expected: false},
		{age: 22, expected: false},
	}

	for _, test := range testParameters {
		name := fmt.Sprintf("%d", test.age)
		t.Run(name, func(t *testing.T) {
			result := conditionals1.IsOldEnoughToDrinkAndDrive(test.age)
			if result != test.expected {
				t.Errorf("%t != %t", result, test.expected)
			}

		})
	}
}
