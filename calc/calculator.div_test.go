//go:build unit
// +build unit

package calc

import (
	"testing"
)

type DivResult struct {
	x        float64
	y        float64
	expected float64
}

var (
	DivResults = []DivResult{
		{100, 10, 10},
		{5645600, 55, 102647.27272727272},
		{2900, 40, 72.5},
		{100, 87, 1.1494252873563218},
	}
	div_c DivCommand
)

func TestDivision(t *testing.T) {
	for _, test := range DivResults {
		div_c.num1 = test.x
		div_c.num2 = test.y
		// fmt.Println(add_c.num1, add_c.num2)
		result, _ := div_c.Run()
		if result != test.expected {
			t.Fatal("Result don't match expected")
		}
	}
}
