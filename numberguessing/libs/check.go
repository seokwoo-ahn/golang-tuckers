package libs

import "fmt"

func Check(input int, answer int) bool {
	if input == answer {
		fmt.Println("정답입니다")
		return true
	} else if input > answer {
		fmt.Println("정답은 더 작은 값입니다")
		return false
	} else if input < answer {
		fmt.Println("정답은 더 큰 값입니다")
		return false
	} else {
		fmt.Println("예기치 못한 error 발생")
		return false
	}
}
