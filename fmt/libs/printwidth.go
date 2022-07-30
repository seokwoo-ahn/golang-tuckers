package libs

import "fmt"

func PrintWidth() {
	var a = 123
	var b = 456
	var c = 123456789

	fmt.Printf("%5d, %5d\n", a, b)     // width: 5
	fmt.Printf("%05d, %05d\n", a, b)   // width: 5, empty space: 0
	fmt.Printf("%-05d, %-05d\n", a, b) // left

	fmt.Printf("%5d, %5d\n", c, c) // shorter than value
	fmt.Printf("%05d, %05d\n", c, c)
	fmt.Printf("%-05d, %-05d\n", c, c)
}
