package main

import (
	"fmt"
	"io"
	"os"
	"strconv"

	"calc-lib/calc"
)

type Calculator interface {
	Calculate(a, b int) int
}

func main() {
	var (
		input      []string   = os.Args[1:]
		output     io.Writer  = os.Stdout
		calculator Calculator = calc.Addition{}
	)
	if len(input) != 2 {
		panic("Usage: <a> <b>")
	}

	a, err := strconv.Atoi(input[0])
	if err != nil {
		panic(err)
	}
	b, err := strconv.Atoi(input[1])
	if err != nil {
		panic(err)
	}

	c := calculator.Calculate(a, b)
	_, err = fmt.Fprintf(output, "%d", c)
	if err != nil {
		panic(err)
	}

}
