package libs

import "fmt"

func AddIndex() {
	slice1 := []int{1, 2, 3, 4, 5, 6}

	slice1 = append(slice1, 0)

	idx := 2

	for i := len(slice1) - 2; i >= idx; i-- {
		slice1[i+1] = slice1[i]
	}

	slice1[idx] = 100

	fmt.Println(slice1)

	slice2 := []int{7, 8, 9, 10, 11}

	slice2 = append(slice2, 0)
	copy(slice2[idx+1:], slice2[idx:])
	slice2[idx] = 100

	fmt.Println(slice2)
}
