package main

import (
	"fmt"
	"golang-tuckers/const/libs"
)

func main() {
	// constant
	fmt.Println("Constant property")
	libs.Constant()
	fmt.Println("Cannot assign & take address to constant")
	// iota
	fmt.Println("Iota")
	libs.Iota()
}
