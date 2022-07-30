package libs

import "fmt"

func Print() {
	var a int = 10
	var b int = 20
	var f float64 = 32799438743.8297

	fmt.Print("a: ", a, "b: ", b, "f: ", f, " end of Print")
	fmt.Println("a: ", a, "b: ", b, "f: ", f, " end of Println")
	fmt.Printf("a: %d b: %d f: %f\n end of Printf", a, b, f)
}
