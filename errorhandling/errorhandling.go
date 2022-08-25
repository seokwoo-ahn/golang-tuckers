package main

import (
	"fmt"
	"golang-tuckers/errorhandling/libs"
)

func main() {
	// readfile with error
	fmt.Println("read or write file")
	libs.ReadFile("test.txt")
	// custom error
	fmt.Println("custom error")
	libs.CustomError()
	// error interface
	fmt.Println("error interface")
	libs.ErrorInterface("cranberry", "izone")
}
