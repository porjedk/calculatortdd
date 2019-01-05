package main

import (
"fmt"
calc "github.com/porjedk/calculatortdd/calculator"
)

func main() {
    //reading input string
    var input string
    fmt.Println("Please enter two numbers in format 'x,x'")
	fmt.Scan(&input)
	fmt.Println("Sum of your input is ", calc.Calculator(input))
}