package libs

import (
	"fmt"
	"sync"
	"time"
)

func waitInt(wg *sync.WaitGroup, ch chan int) {
	for n := range ch {
		fmt.Println("input int:", n)
		time.Sleep(time.Second)
	}
	wg.Done()
}

func Wait() {
	var wg sync.WaitGroup
	ch := make(chan int)

	wg.Add(1)
	go waitInt(&wg, ch)

	for i := 0; i < 10; i++ {
		ch <- i * 2
	}
	close(ch)
	wg.Wait()
}
