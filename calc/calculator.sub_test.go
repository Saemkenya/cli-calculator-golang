//go:build unit
// +build unit

package calc

import (
	"testing"
)

type SubResult struct {
	x        float64
	y        float64
	expected float64
}

var (
	SubResults = []SubResult{
		{92304, 41234, 51070},
		{219475, 237642, -18167},
		{-64723, 9563, -74286},
		{748322, 34253, 714069},
		{897425, -6742386, 7639811},
	}
	sub_c SubCommand
)

func TestSubtraction(t *testing.T) {
	for _, test := range SubResults {
		sub_c.num1 = test.x
		sub_c.num2 = test.y
		result, _ := sub_c.Run()
		if result != test.expected {
			t.Fatal("Result don't match expected")
		}
	}
}
