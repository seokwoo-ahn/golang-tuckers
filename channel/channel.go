package main

import (
	"fmt"
	"golang-tuckers/channel/libs"
)

func main() {
	// data channel
	fmt.Println("channel gen")
	libs.Data()
	// channel size
	fmt.Println("channel size")
	libs.Size(10)
	// channel wait
	fmt.Println("channel wait")
	libs.Wait()
}
