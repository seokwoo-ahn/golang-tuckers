package libs

import "fmt"

func Overflow() {
	var a int16 = 3456
	var c int8 = int8(a) // int overflow

	fmt.Println(a)
	fmt.Println(c)
}
