//go:build unit
// +build unit

package calc

import (
	"testing"
)

type AddResult struct {
	x        float64
	y        float64
	expected float64
}

var (
	AddResults = []AddResult{
		{1, 1, 2},
		{10, -10, 0},
		{87, 100, 187},
	}
	add_c AddCommand
)

func TestAddition(t *testing.T) {
	for _, test := range AddResults {
		add_c.num1 = test.x
		add_c.num2 = test.y
		result, _ := add_c.Run()
		if result != test.expected {
			t.Fatal("Result don't match expected")
		}
	}
}
