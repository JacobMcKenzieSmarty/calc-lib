package main

import (
	"log"
	"os"

	"github.com/JacobMcKenzieSmarty/calc-lib/calc"
	"github.com/JacobMcKenzieSmarty/calc-lib/handler"
)

func main() {
	err := handler.NewHandler(os.Stdout, calc.Addition{}).Handle(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
}
