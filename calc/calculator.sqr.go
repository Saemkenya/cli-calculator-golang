package calc

import (
	"flag"
	"fmt"
	"math"
)

func NewSquareCommand() *SquareCommand {
	cc := &SquareCommand{
		fs: flag.NewFlagSet("sqr", flag.ContinueOnError),
	}

	cc.fs.Float64Var(&cc.num1, "num1", 100, "number to find the square root of")
	cc.fs.Float64Var(&cc.num2, "num2", 0, "second number not required for square root func")
	cc.fs.BoolVar(&cc.show, "show", false, "square root formula details")

	return cc
}

type SquareCommand struct {
	fs *flag.FlagSet

	num1 float64
	num2 float64
	show bool
}

func (sc *SquareCommand) Name() string {
	fmt.Println("Square root")
	return sc.fs.Name()
}

func (sc *SquareCommand) Print() {
	sc.fs.PrintDefaults()
	return
}

func (sc *SquareCommand) Init(args []string) error {
	return sc.fs.Parse(args)
}

func (sc *SquareCommand) Run() (result float64, err error) {
	err = nil
	result = math.Sqrt(sc.num1)
	fmt.Println(result)
	if sc.num2 > 0 {
		fmt.Println("Num 2 arg is not required in square-root func")
	} else if sc.show == true {
		fmt.Printf("Default formula used to find square-root of %v\n", sc.num1)
		fmt.Printf("Result %f\n", result)
	} else {
		fmt.Printf("Result %f\n", result)
	}
	return
}
