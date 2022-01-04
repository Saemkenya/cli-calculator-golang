//go:build unit
// +build unit

package calc

import (
	"testing"
)

type MultResult struct {
	x        float64
	y        float64
	expected float64
}

var (
	MultResults = []MultResult{
		{32424, 95234, 3087867216},
		{56234756, 94523, 5315477841388},
		{863467, 563425, 486498894475},
		{-5634786, 54623, -307788915678},
		{4812673, -5463, -26291632599},
	}
	mult_c MultCommand
)

func TestMultiplication(t *testing.T) {
	for _, test := range MultResults {
		mult_c.num1 = test.x
		mult_c.num2 = test.y
		result, _ := mult_c.Run()
		if result != test.expected {
			t.Fatal("Result don't match expected")
		}
	}
}
