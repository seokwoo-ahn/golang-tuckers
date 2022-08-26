package libs

import (
	"fmt"
	"sync"
)

var Wg sync.WaitGroup

func WaitGroup(a, b int) {
	sum := 0
	for i := a; i <= b; i++ {
		sum += i
	}
	fmt.Println("Sum:", sum)
	Wg.Done()
}
