package libs

import "fmt"

func Append() {
	var slice = make([]int, 0)
	slice = append(slice, 1, 2, 3, 4, 5)
	fmt.Println(slice)
}
