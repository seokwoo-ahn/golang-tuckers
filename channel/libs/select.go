package libs

import (
	"fmt"
	"sync"
	"time"
)

func test(wg *sync.WaitGroup, ch chan int, quit chan bool) {
	for {
		select {
		case n := <-ch:
			fmt.Println("input int is:", n)
			time.Sleep(time.Second)
		case <-quit:
			wg.Done()
			return
		}
	}
}
func Select() {
	var wg sync.WaitGroup
	ch := make(chan int)
	quit := make(chan bool)

	wg.Add(1)
	go test(&wg, ch, quit)

	for i := 0; i < 10; i++ {
		ch <- i * 2
	}

	quit <- true
	wg.Wait()
}
