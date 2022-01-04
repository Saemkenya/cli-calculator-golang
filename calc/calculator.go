package calc

type ICalculator interface {
	Init([]string) error
	Run() (float64, error)
	Name() string
	Print()
}
