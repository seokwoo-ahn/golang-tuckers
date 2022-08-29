package libs

import (
	"fmt"
	"sync"
	"time"
)

func check(wg *sync.WaitGroup, ch chan int) {
	tick := time.Tick(time.Second)
	terminate := time.After(5 * time.Second)
	// ch 보다 빠르게 종료되면 영원히 종료되지 않음

	for {
		select {
		case <-tick:
			fmt.Println("tick")
		case <-terminate:
			fmt.Println("terminated!")
			wg.Done()
			return
		case n := <-ch:
			fmt.Println("input:", n)
			time.Sleep(time.Microsecond.Abs())
		}
	}
}

func Interval() {
	var wg sync.WaitGroup
	ch := make(chan int)
	wg.Add(1)
	go check(&wg, ch)

	for i := 0; i < 10; i++ {
		ch <- i * 2
	}
	wg.Wait()
}
