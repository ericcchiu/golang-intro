package lab1

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var step = 1

var numbers = [5]int{}

func TestBasicArray(t *testing.T) {
	BasicArray()
	if step >= 1 {
		for i, v := range BasicArray() {
			if ok := assert.Equal(t, v, i+1, fmt.Sprintf("%d should equal %d", v, i+1)); !ok {
				return
			}
			step = 2
		}
	}
}

var food = []string{}

func TestBasicSlice(t *testing.T) {
	expect := []string{"apple", "pizza", "taco", "cheese"}
	if step >= 2 {
		if ok := assert.Equal(t, BasicSlice(), expect, "Slice should contain the correct foods"); !ok {
			return
		}
		step = 3
	} else {

		t.Skip(fmt.Sprintf("Not ready for test %d", 2))
	}
}

func TestMergeAndSubSlice(t *testing.T) {

	tests := []struct {
		morefood   []string
		expectfull []string
		expectsub  []string
	}{
		{
			morefood:   []string{"bacon", "eggs", "cheese"},
			expectfull: []string{"apple", "pizza", "bacon", "eggs", "cheese"},
			expectsub:  []string{"pizza", "bacon", "eggs"},
		},
		{
			morefood:   []string{"chicken", "olives", "toast", "ham sammich"},
			expectfull: []string{"apple", "pizza", "chicken", "olives", "toast", "ham sammich"},
			expectsub:  []string{"pizza", "chicken", "olives"},
		},
	}

	if step >= 3 {

		for _, test := range tests {
			f, s := MergeAndSubSlice(test.morefood)
			if ok := assert.Equal(t, test.expectfull, f); !ok {
				return
			}
			if ok := assert.Equal(t, test.expectsub, s); !ok {
				return
			}
		}

		step = 4
	} else {

		t.Skip(fmt.Sprintf("Not ready for test %d", 3))
	}
}
