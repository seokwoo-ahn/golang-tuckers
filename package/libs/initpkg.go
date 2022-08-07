package libs

import "fmt"

var (
	a = c + b
	b = f()
	c = f()
	d = 3
)

func init() {
	d++
	fmt.Println("init function", d)
	fmt.Println("initpkg package init complete")
}

func f() int {
	d++
	fmt.Println("f() d:", d)
	return d
}
