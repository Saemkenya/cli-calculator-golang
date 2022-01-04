package calc

import (
	"flag"
	"fmt"
)

func NewDivCommand() *DivCommand {
	cc := &DivCommand{
		fs: flag.NewFlagSet("div", flag.ContinueOnError),
	}

	cc.fs.Float64Var(&cc.num1, "num1", 100, "first number of the division")
	cc.fs.Float64Var(&cc.num2, "num2", 87, "second number of the division")
	cc.fs.BoolVar(&cc.show, "show", false, "division formula details")

	return cc
}

type DivCommand struct {
	fs *flag.FlagSet

	num1 float64
	num2 float64
	show bool
}

func (dc *DivCommand) Name() string {
	fmt.Println("Division")
	return dc.fs.Name()
}

func (dc *DivCommand) Print() {
	dc.fs.PrintDefaults()
	return
}

func (dc *DivCommand) Init(args []string) error {
	return dc.fs.Parse(args)
}

func (dc *DivCommand) Run() (result float64, err error) {
	// fmt.Println(dc.num1, dc.num2)
	err = nil
	result = dc.num1 / dc.num2
	if dc.num2 == 0 {
		fmt.Println("Can't perform a division by zero(0)")
		return
	} else if dc.show == true {
		fmt.Printf("Default formula used to divide %f and %f\n", dc.num1, dc.num2)
		fmt.Printf("Result %f\n", result)
		return
	} else {
		fmt.Printf("Result %f\n", result)
	}
	return
}
