package libs

import "fmt"

func Defer() {
	defer fmt.Println("가장 첫번쨰로 선언된 defer 입니다")
	defer fmt.Println("두번째로 선언된 defer 입니다")

	fmt.Println("마지막으로 작성된 로직입니다")
}
