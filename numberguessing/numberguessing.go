package main

import (
	"fmt"
	"golang-tuckers/numberguessing/libs"
)

func main() {
	var answer = libs.NumGen()
	input, err := libs.Scan()
	if err != nil {
		fmt.Println("적절한 숫자가 아닙니다! 다시 입력하세요")
	}

	fmt.Println(answer, input)
}
