package main

import (
	"flag"
	"log"
	"os"

	"github.com/JacobMcKenzieSmarty/calc-lib/calc"
	"github.com/JacobMcKenzieSmarty/calc-lib/handler"
)

func main() {
	var op string
	flag.StringVar(&op, "op", "+", "One of + - * /")
	flag.Parse()
	err := handler.NewHandler(os.Stdout, calculators[op]).Handle(flag.Args())
	if err != nil {
		log.Fatal(err)
	}
}

var calculators = map[string]calc.Calculator{
	"+": calc.Addition{},
	"-": calc.Subtraction{},
	"*": calc.Multiplication{},
	"/": calc.Division{},
}
