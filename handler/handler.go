package handler

import (
	"errors"
	"fmt"
	"io"
	"strconv"

	"github.com/JacobMcKenzieSmarty/calc-lib/calc"
)

type Handler struct {
	calculator calc.Calculator
	output     io.Writer //lets you test in memory now (to a buffer, etc.) - an abstraction, not a low level detail.
}

func (this *Handler) Handle(args []string) error {
	if len(args) != 2 {
		return fmt.Errorf("%w: two args req (you provided %d", ErrTooFewArgs, len(args))
	}

	a, err := strconv.Atoi(args[0])
	if err != nil {
		return fmt.Errorf("%w: first arg (%s) %w", ErrMalformedArgs, a, err)
	}
	b, err := strconv.Atoi(args[1])
	if err != nil {
		return fmt.Errorf("%w: first arg (%s) %w", ErrMalformedArgs, b, err)
	}

	c := this.calculator.Calculate(a, b)
	_, err = fmt.Fprintf(this.output, "%v\n", c)
	if err != nil {
		return fmt.Errorf("%w, %w", ErrOutputWriter, err)
	}

	return nil
}

func NewHandler(output io.Writer, calculator calc.Addition) *Handler {
	return &Handler{
		calculator: calculator,
		output:     output,
	}
}

var (
	ErrTooFewArgs    = errors.New("\"Usage: <a> <b>\"")
	ErrMalformedArgs = errors.New("invalid argument")
	ErrOutputWriter  = errors.New("output write failed")
)
