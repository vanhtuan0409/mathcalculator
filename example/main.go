package main

import (
	"fmt"

	"github.com/vanhtuan0409/mathcalculator"
)

func main() {
	expression := "1+2    * (3 + 4) - 5"
	result, err := mathcalculator.Evaluate(expression)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(expression, "=", result)
	}
}
