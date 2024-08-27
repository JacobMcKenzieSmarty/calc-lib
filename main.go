package main

import (
	"log"
	"os"

	"calc-lib/calc"
	"calc-lib/handler"
)

func main() {
	err := handler.NewHandler(os.Stdout, calc.Addition{}).Handle(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
}
