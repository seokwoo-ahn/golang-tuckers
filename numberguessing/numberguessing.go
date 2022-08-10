package main

import (
	"fmt"
	"golang-tuckers/numberguessing/libs"
)

func main() {
	var answer = libs.NumGen()
	var flag = false
	var cnt = 0

	for flag == false {
		cnt++
		input, err := libs.Scan()
		if err != nil {
			fmt.Println("적절한 숫자가 아닙니다! 다시 입력하세요")
			continue
		}
		flag = libs.Check(input, answer)
	}

	fmt.Println("축하합니다! 시도한 횟수:", cnt)
}
