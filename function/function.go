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
}
