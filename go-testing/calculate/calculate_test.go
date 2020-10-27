package calculate

import (
	"testing"

	"github.com/ericcchiu/golang-intro/go-testing/calculate"
)

func TestAdd(t *testing.T) {
	result := calculate.Add(3, 2)
	if result != 5 {
		t.Errorf("Expected 5, got %v", result)
	}
}
