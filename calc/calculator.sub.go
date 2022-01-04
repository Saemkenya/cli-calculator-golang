package calc

import (
	"flag"
	"fmt"
)

func NewSubCommand() *SubCommand {
	cc := &SubCommand{
		fs: flag.NewFlagSet("sub", flag.ContinueOnError),
	}

	cc.fs.Float64Var(&cc.num1, "num1", 100, "first number of the subtraction")
	cc.fs.Float64Var(&cc.num2, "num2", 87, "second number of the subtraction")
	cc.fs.BoolVar(&cc.show, "show", false, "subtraction formula details")

	return cc
}

type SubCommand struct {
	fs *flag.FlagSet

	num1 float64
	num2 float64
	show bool
}

func (sc *SubCommand) Name() string {
	fmt.Println("Subtraction")
	return sc.fs.Name()
}

func (sc *SubCommand) Print() {
	sc.fs.PrintDefaults()
	return
}

func (sc *SubCommand) Init(args []string) error {
	return sc.fs.Parse(args)
}

func (sc *SubCommand) Run() (result float64, err error) {
	err = nil
	result = sc.num1 - sc.num2
	if sc.show == true {
		fmt.Printf("Default formula used to subtract %f and %f\n", sc.num1, sc.num2)
		fmt.Printf("Result %f\n", result)
	} else {
		fmt.Printf("Result %f\n", result)
	}
	return
}
