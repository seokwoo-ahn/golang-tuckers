package libs

import "fmt"

func FallThrough(num int) {

	a := num

	switch a {
	case 1:
		fmt.Println("a == 1")
		fallthrough
	case 2:
		fmt.Println("a == 2")
		fallthrough
	case 3:
		fmt.Println("a == 3")
	case 4:
		fmt.Println("a == 4")
		fallthrough
	case 5:
		fmt.Println("a == 5")
		fallthrough
	default:
		fmt.Println("a > 4")
	}
}
