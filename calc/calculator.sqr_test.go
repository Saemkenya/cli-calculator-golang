//go:build unit
// +build unit

package calc

import (
	"testing"
)

type SqrResult struct {
	x        float64
	expected float64
}

var (
	SqrResults = []SqrResult{
		{534673, 731.2133751511934},
		{789242, 888.3929310839883},
		{453217, 673.213933307979},
		{67423, 259.65939228150404},
		{122563, 350.089988431546},
	}
	sqr_c SquareCommand
)

func TestSquareroot(t *testing.T) {
	for _, test := range SqrResults {
		sqr_c.num1 = test.x
		result, _ := sqr_c.Run()
		if result != test.expected {
			t.Fatal("Result don't match expected")
		}
	}
}
