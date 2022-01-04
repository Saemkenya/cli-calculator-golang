package calc

import (
	"flag"
	"fmt"
)

func NewMultCommand() *MultCommand {
	cc := &MultCommand{
		fs: flag.NewFlagSet("mult", flag.ContinueOnError),
	}

	cc.fs.Float64Var(&cc.num1, "num1", 100, "first number of the multiplication")
	cc.fs.Float64Var(&cc.num2, "num2", 87, "second number of the multiplication")
	cc.fs.BoolVar(&cc.show, "show", false, "multiplication formula details")

	return cc
}

type MultCommand struct {
	fs *flag.FlagSet

	num1 float64
	num2 float64
	show bool
}

func (mc *MultCommand) Name() string {
	fmt.Println("Multiplication")
	return mc.fs.Name()
}

func (mc *MultCommand) Print() {
	mc.fs.PrintDefaults()
	return
}

func (mc *MultCommand) Init(args []string) error {
	return mc.fs.Parse(args)
}

func (mc *MultCommand) Run() (result float64, err error) {
	err = nil
	result = mc.num1 * mc.num2
	if mc.show == true {
		fmt.Printf("Default formula used to multiply %f and %f\n", mc.num1, mc.num2)
		fmt.Printf("Result %f\n", result)
	} else {
		fmt.Printf("Result %f\n", result)
	}
	return
}
