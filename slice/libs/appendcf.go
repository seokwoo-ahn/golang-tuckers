package libs

import "fmt"

func AppendCf() {
	fmt.Println("not enough capacity")
	slice1 := []int{1, 2, 3}
	slice2 := append(slice1, 4, 5)

	fmt.Println("slice1:", slice1, len(slice1), cap(slice1))
	fmt.Println("slice2:", slice2, len(slice2), cap(slice2))

	slice1[1] = 100

	fmt.Println("After change second element")

	fmt.Println("slice1:", slice1, len(slice1), cap(slice1))
	fmt.Println("slice2:", slice2, len(slice2), cap(slice2))

	slice1 = append(slice1, 500)

	fmt.Println("After append 500")
	fmt.Println("slice1:", slice1, len(slice1), cap(slice1))
	fmt.Println("slice2:", slice2, len(slice2), cap(slice2))

	fmt.Println("enough capacity")
	slice3 := make([]int, 3, 5)
	slice4 := append(slice3, 4, 5)

	fmt.Println("slice3:", slice3, len(slice3), cap(slice3))
	fmt.Println("slice4:", slice4, len(slice4), cap(slice4))

	slice3[1] = 100

	fmt.Println("slice3:", slice3, len(slice3), cap(slice3))
	fmt.Println("slice4:", slice4, len(slice4), cap(slice4))

	slice3 = append(slice3, 500)

	fmt.Println("After append 500")
	fmt.Println("slice3:", slice3, len(slice3), cap(slice3))
	fmt.Println("slice4:", slice4, len(slice4), cap(slice4))
}
