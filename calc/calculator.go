package calc

type ICalculator interface {
	Init([]string) error
	Run() error
	Name() string
	Print()
}
