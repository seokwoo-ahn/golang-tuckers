package libs

import "fmt"

func ArrayCopy() {
	a := [5]int{1, 2, 3, 4, 5}
	b := [5]int{6, 7, 8, 9, 10}

	a = b

	for _, v := range a {
		fmt.Println(v)
	}
}
