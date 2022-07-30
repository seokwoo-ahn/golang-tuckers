package libs

import (
	"bufio"
	"fmt"
	"os"
)

func Scan() {
	stdin := bufio.NewReader(os.Stdin)

	var a int
	var b int

	var c int
	var d int

	var e int
	var f int

	//scan
	n1, err1 := fmt.Scan(&a, &b)

	if err1 != nil {
		stdin.ReadString('\n')
	}

	// scanf
	n2, err2 := fmt.Scanf("%d %d\n", &c, &d)

	if err2 != nil {
		stdin.ReadString('\n')
	}

	// scanln
	n3, err3 := fmt.Scanln(&e, &f)

	if err3 != nil {
		stdin.ReadString('\n')
	}

	fmt.Println(n1, a, b, err1)
	fmt.Println(n2, c, d, err2)
	fmt.Println(n3, e, f, err3)
}
