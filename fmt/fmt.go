package main

import (
	"fmt"
	"golang-tuckers/fmt/libs"
)

func main() {
	// print test
	fmt.Println("fmt print test")
	libs.Print()
	// width test
	fmt.Println("print width test")
	libs.PrintWidth()
	// special symbols test
	fmt.Println("special sysmbols test")
	libs.SpecialSymbols()
	// scan test
	fmt.Println("scan test")
	libs.Scan()
}
