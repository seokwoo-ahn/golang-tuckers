package main

import (
	"fmt"
	"golang-tuckers/interface/libs"
)

func main() {
	// interface declaration
	student := libs.Student{Name: "cranberry", Age: 12}
	var stringer libs.Stringer = student
	fmt.Println(stringer.Declaration())
	// empty interface
	fmt.Println("empty interface")
	libs.EmptyInterface("hello")
	libs.EmptyInterface(123)
	// interface convert to concrete type
	fmt.Println("interface convert")
	str := &libs.Str{Name: "cranberry", Age: 12}
	libs.PrintStr(str)
	// interface convert to interface
	fmt.Println("interface convert to interface")
	file := &libs.File{}
	libs.InterfaceConvert(file)
}
