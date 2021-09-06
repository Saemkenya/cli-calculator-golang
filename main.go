package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/Saemkenya/cli-calculator-golang/calc"
)

func check(args []string) error {

	if len(args) < 1 {
		return errors.New("A subcommand is necessary here chief ...")
	}
	cmds := []calc.ICalculator{
		calc.NewAddCommand(),
		calc.NewSubCommand(),
		calc.NewMultCommand(),
		calc.NewDivCommand(),
		calc.NewSquareCommand(),
	}

	if len(args) < 1 {
		for _, cmd := range cmds {
			cmd.Name()
			cmd.Print()
		}
		return errors.New("You must pass a sub-command")
	}
	return nil
}

func root(args []string, _command calc.ICalculator) error {
	cmds := []calc.ICalculator{
		_command,
	}

	subcommand := os.Args[1]

	for _, cmd := range cmds {
		if cmd.Name() == subcommand {
			cmd.Init(os.Args[2:])
			return cmd.Run()
		}
	}
	return fmt.Errorf("Unknown subcommand: %s", subcommand)
}

func doCalc(arg string, args []string) {
	switch arg {
	case "add":
		if err := root(args, calc.NewAddCommand()); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	case "sub":
		if err := root(args, calc.NewSubCommand()); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	case "mult":
		if err := root(args, calc.NewMultCommand()); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	case "div":
		if err := root(args, calc.NewDivCommand()); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	case "sqr":
		if err := root(args, calc.NewSquareCommand()); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	default:
		fmt.Printf("unknown calculator sub-command: %v.\n", arg)
	}
}

func main() {
	var args []string = os.Args[1:] // os.Args[1:]
	if err := check(args); err != nil {
		fmt.Println(err.Error())
		return
	}

	var arg string = os.Args[1]
	doCalc(arg, args)
}
