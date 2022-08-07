package main

import (
	"fmt"
	"golang-tuckers/package/libs"
)

func main() {
	// import package with path
	fmt.Println("import package with path")
	libs.Path()
	// alias
	fmt.Println("alias template")
	libs.Alias()
	// unused package
	fmt.Println("unused package")
	fmt.Println(libs.Unused("cranberry"))
	// external package
	fmt.Println("external package")
	libs.ExternalPkg()
}
