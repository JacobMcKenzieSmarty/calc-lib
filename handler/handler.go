package handler

import (
	"fmt"
	"io"
	"strconv"

	"calc-lib/calc"
)

type Handler struct {
	input      []string
	output     io.Writer
	calculator calc.Calculator
}

func (this *Handler) Handle() {
	if len(this.input) != 2 {
		panic("Usage: <a> <b>")
	}

	a, err := strconv.Atoi(this.input[0])
	if err != nil {
		panic(err)
	}
	b, err := strconv.Atoi(this.input[1])
	if err != nil {
		panic(err)
	}

	c := this.calculator.Calculate(a, b)
	_, err = fmt.Fprintf(this.output, "%v\n", c)
	if err != nil {
		panic(err)
	}
}

func NewHandler(input []string, output io.Writer, calculator calc.Calculator) *Handler {
	return &Handler{
		input:      input,
		output:     output,
		calculator: calculator,
	}
}
