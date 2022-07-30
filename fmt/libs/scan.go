package libs

import "fmt"

func Scan() {
	var a int
	var b int

	var c int
	var d int

	var e int
	var f int

	//scan
	n1, err1 := fmt.Scan(&a, &b)

	// scanf
	n2, err2 := fmt.Scanf("%d %d\n", &c, &d)

	// scanln
	n3, err3 := fmt.Scanln(&e, &f)

	fmt.Println(n1, a, b, err1)
	fmt.Println(n2, c, d, err2)
	fmt.Println(n3, e, f, err3)
}
