package main

import (
	"fmt"
	"golang-tuckers/struct/libs"
)

func main() {
	// struct declaration
	fmt.Println("struct declaration")
	libs.Declaration()
	// embedded struct
	fmt.Println("embedded struct")
	libs.Embedded()
	// struct copy
	fmt.Println("struct copy")
	libs.Copy()
}
