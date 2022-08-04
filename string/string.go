package main

import (
	"fmt"
	"golang-tuckers/string/libs"
)

func main() {
	// string declaration
	fmt.Println("string declaration")
	libs.Declaration()
	// string length
	fmt.Println("string length")
	libs.Length()
	// number of charaters
	fmt.Println("string count")
	libs.StringCount("안녕하세요 반갑습니다")
	// range
	fmt.Println("string range")
	libs.Range("아이즈원")
	// combine
	fmt.Println("append string")
	fmt.Println(libs.Combine("Cran", "Berry"))
	//	compare
	fmt.Println("compare string")
	fmt.Println(libs.Compare("berry", "berry"))
	// duplication
	fmt.Println("duplication not value but pointer")
	libs.Duplication()
	// string change
	fmt.Println("change string")
	libs.Change()
}
