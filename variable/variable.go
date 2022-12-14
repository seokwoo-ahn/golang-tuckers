package main

import (
	"fmt"
	"golang-tuckers/variable/libs"
)

func main() {
	//variable test
	fmt.Println("sample test")
	libs.Sample()
	//declaration test
	fmt.Println("declaration test")
	libs.Declaration()
	//type test
	fmt.Println("type test")
	libs.Type()
	//int overflow test
	fmt.Println("int overflow test")
	libs.Overflow()
	//scope test
	fmt.Println("scope test")
	libs.Scope()
	//float error test
	fmt.Println("float error test")
	libs.Floaterror()
}
