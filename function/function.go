package main

import (
	"fmt"
	"golang-tuckers/function/libs"
)

func main() {
	// add function
	fmt.Println("add function")
	add := libs.Add(2, 3)
	fmt.Println("2 + 3 =", add)
	// average function
	fmt.Println("average function")
	libs.Average("seok", 60, 70, 80)
	// multi return (max value)
	fmt.Println("max function")
	max, result := libs.MultiReturn(2.44, 5.11)
	fmt.Println(max, result)
	// recursive call
	fmt.Println("reculsive call")
	libs.PrintNumSequence(5)
}
