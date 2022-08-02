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
}
