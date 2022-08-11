package libs

import "fmt"

func Replicate() {
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := make([]int, len(slice1))

	for i, v := range slice1 {
		slice2[i] = v
	}

	slice1[1] = 100
	fmt.Println(slice1)
	fmt.Println(slice2)

	slice3 := []int{6, 7, 8, 9, 10}
	slice4 := append([]int{}, slice3...)

	slice3[1] = 600
	fmt.Println(slice3)
	fmt.Println(slice4)

	fmt.Println("copy slice")
	slice5 := []int{11, 12, 13, 14, 15}
	slice6 := make([]int, 3, 10)
	slice7 := make([]int, 10)

	cnt1 := copy(slice6, slice5)
	cnt2 := copy(slice7, slice5)

	fmt.Println(cnt1, slice6)
	fmt.Println(cnt2, slice7)
}
