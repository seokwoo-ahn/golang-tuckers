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
	// replicate slice
	fmt.Println("replicate slice")
	libs.Replicate()
	// delete index
	fmt.Println("delete index")
	libs.DeleteIndex()
	// add index
	fmt.Println("add index")
	libs.AddIndex()
	// slice sort
	fmt.Println("slice sort")
	libs.Sort()
}
