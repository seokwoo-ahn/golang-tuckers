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

	var d = 324.13455
	var e = 3.14

	fmt.Printf("%08.2f\n", d) // width: 8, empty space:0, decimal: 2
	fmt.Printf("%08.2g\n", d) // width: 8, empty space:0, total num: 2
	fmt.Printf("%8.5g\n", d)  // width: 8, totla num: 5
	fmt.Printf("%f\n", e)     // default deciaml: 6
}
