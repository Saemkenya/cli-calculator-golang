package calc

import (
	"flag"
	"fmt"
)

func NewAddCommand() *AddCommand {
	cc := &AddCommand{
		fs: flag.NewFlagSet("add", flag.ContinueOnError),
	}

	cc.fs.Float64Var(&cc.num1, "num1", 100, "first number of the addition")
	cc.fs.Float64Var(&cc.num2, "num2", 87, "second number of the addition")
	cc.fs.BoolVar(&cc.show, "show", false, "addition formula details")

	return cc
}

type AddCommand struct {
	fs *flag.FlagSet

	num1 float64
	num2 float64
	show bool
}

func (ac *AddCommand) Name() string {
	fmt.Println("Addition")
	return ac.fs.Name()
}

func (ac *AddCommand) Print() {
	ac.fs.PrintDefaults()
	return
}

func (ac *AddCommand) Init(args []string) error {
	return ac.fs.Parse(args)
}

func (ac *AddCommand) Run() error {
	if ac.show == true {
		fmt.Printf("Default formula used to add %f and %f\n", ac.num1, ac.num2)
		fmt.Printf("Result %f\n", ac.num1+ac.num2)
	} else {
		fmt.Printf("Result %f\n", ac.num1+ac.num2)
	}
	return nil
}
