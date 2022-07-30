package libs

import "fmt"

func Type() {
	a := 3              // int
	var b float64 = 3.5 // float64

	var c int = int(b)  // float64 => int
	d := float64(a * c) // int => float64

	var e int64 = 7
	f := int64(d) * e // float64 => int64

	var g int = int(b * 3)
	var h int = int(b) * 3 // int(b) = 3
	fmt.Println(g, h, f)
}
