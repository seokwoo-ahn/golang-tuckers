package libs

import "fmt"

var a int = 10 // 패키지 전역 변수

func Scope() {
	var b int = 20

	{
		var c int = 50
		fmt.Println(a, b, c)
	}

	// b = c + 30 => error
}
