package libs

import "fmt"

func PrintNumSequence(num int) {
	if num == 0 {
		return
	}
	fmt.Println(num)
	PrintNumSequence(num - 1)
	fmt.Println("call num:", num)
}
