package handler

import (
	"errors"
	"fmt"
	"io"
	"strconv"

	"calc-lib/calc"
)

type Handler struct {
	calculator calc.Calculator
	output     io.Writer //lets you test in memory now (to a buffer, etc.) - an abstraction, not a low level detail.
}

func (this *Handler) Handle(args []string) error {
	if len(args) != 2 {
		return errors.New("Usage: <a> <b>")
	}

	a, err := strconv.Atoi(args[0])
	if err != nil {
		return err
	}
	b, err := strconv.Atoi(args[1])
	if err != nil {
		return err
	}

	c := this.calculator.Calculate(a, b)
	_, err = fmt.Fprintf(this.output, "%v\n", c)
	if err != nil {
		return err
	}

	return nil
}

func NewHandler(output io.Writer, calculator calc.Addition) *Handler {
	return &Handler{
		calculator: calculator,
		output:     output,
	}
}
