package main

import (
	"fmt"
	"golang-tuckers/goroutine/libs"
)

func main() {
	// gen go routine
	fmt.Println("go routine")
	libs.GenGoRoutine()
	// wait group
	fmt.Println("wait group")
	libs.Wg.Add(8)
	for i := 0; i < 8; i++ {
		go libs.WaitGroup(1, 450555)
	}
	libs.Wg.Wait()
}
