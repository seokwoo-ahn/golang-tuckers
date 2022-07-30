package main

import (
	"fmt"
	"golang-tuckers/operator/libs"
)

func main() {
	// arithmetic operation test
	fmt.Println("arithmetic operation test")
	libs.Arithmetic()
	// bit operator test
	fmt.Println("bit operator test")
	libs.Bit()
	// shift operator test
	fmt.Println("shift operator test")
	libs.Shift()
	// int overflow test
	fmt.Println("int overflow test")
	libs.IntOverFlow()
	// float test
	fmt.Println("float test")
	libs.Float()
}
