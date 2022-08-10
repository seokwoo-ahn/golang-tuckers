package libs

import "fmt"

func Declaration() {
	var slice1 = make([]int, 3)
	var slice2 = []int{1, 2, 3}

	slice1[0] = 1

	for i, v := range slice1 {
		fmt.Printf("slice1[%d] = %d\n", i, v)
	}

	for i, v := range slice2 {
		fmt.Printf("slice2[%d] = %d\n", i, v)
	}

}
