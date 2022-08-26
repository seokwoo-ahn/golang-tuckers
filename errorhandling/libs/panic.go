package libs

import "fmt"

func PanicDivide(a, b int) {
	if b == 0 {
		panic("b는 0일 수 없습니다")
	}
	fmt.Printf("%d / %d = %d\n", a, b, a/b)
}
