package main

import (
	"fmt"
	"golang-tuckers/slice/libs"
)

func main() {
	// slice declaration
	fmt.Println("slice declaration")
	libs.Declaration()
	// slice append
	fmt.Println("slice append")
	libs.Append()
	// slice append cf
	fmt.Println("slice append cf")
	libs.AppendCf()
}
