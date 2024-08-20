package main

import (
	"os"

	"calc-lib/calc"
	"calc-lib/handler"
)

func main() {
	handler.NewHandler(os.Args[1:], os.Stdout, calc.Addition{}).Handle()
}
