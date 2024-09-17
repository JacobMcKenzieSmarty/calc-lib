package calc

type Calculator interface {
	Calculate(a, b int) int
}

type Addition struct{}
type Subtraction struct{}
type Multiplication struct{}
type Division struct{}

func (Addition) Calculate(a, b int) int {
	return a + b
}
func (Subtraction) Calculate(a, b int) int    { return a - b }
func (Multiplication) Calculate(a, b int) int { return a * b }
func (Division) Calculate(a, b int) int       { return a / b }
