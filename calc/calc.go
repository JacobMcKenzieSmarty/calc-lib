package calc

type Addition struct{}

func (Addition) Calculate(a, b int) int {
	return a + b
}
