package libs

import "fmt"

func Size(input int) {
	ch := make(chan int, 4)
	ch <- input
	result := <-ch
	fmt.Printf("input is %d\n", result)
}
