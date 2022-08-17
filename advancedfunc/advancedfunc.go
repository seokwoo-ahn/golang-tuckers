package main

import (
	"fmt"
	"golang-tuckers/advancedfunc/libs"
)

func main() {
	// variable arguments
	fmt.Println("variable arguments")
	libs.VariableArg(1, 3, 4)
	libs.VariableArg(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	// variable arguments with empty interface
	fmt.Println("variable arguments with empty interface")
	libs.VariableArgInterface(1, "hi", [2]int{1, 2})
	// defer
	fmt.Println("defer")
	libs.Defer()
	// function variable
	fmt.Println("function variable")
	var operator func(int, int) int = libs.FunctionVariable("*")
	var result = operator(3, 4)
	fmt.Println(result)
}
